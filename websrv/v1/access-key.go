// Copyright 2014 Eryx <evorui аt gmаil dοt cοm>, All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	"fmt"
	"time"

	"github.com/hooto/httpsrv"
	"github.com/lessos/lessgo/crypto/idhash"
	"github.com/lessos/lessgo/types"
	iox_utils "github.com/lynkdb/iomix/utils"

	"github.com/hooto/iam/iamapi"
	"github.com/hooto/iam/iamclient"
	"github.com/hooto/iam/store"
)

var (
	ak_limit = 20
)

type AccessKey struct {
	*httpsrv.Controller
	us iamapi.UserSession
}

func (c *AccessKey) Init() int {

	//
	c.us, _ = iamclient.SessionInstance(c.Session)

	if !c.us.IsLogin() {
		c.Response.Out.WriteHeader(401)
		c.RenderJson(types.NewTypeErrorMeta(iamapi.ErrCodeUnauthorized, "Unauthorized"))
		return 1
	}

	return 0
}

func (c AccessKey) EntryAction() {

	var set struct {
		types.TypeMeta
		*iamapi.AccessKey `json:",omitempty"`
	}
	defer c.RenderJson(&set)

	id := c.Params.Get("access_key")
	if id == "" {
		set.Error = types.NewErrorMeta(iamapi.ErrCodeNotFound, "Access Key Not Found")
		return
	}

	var ak iamapi.AccessKey
	if rs := store.Data.NewReader(iamapi.ObjKeyAccessKey(c.us.UserName, id)).Query(); rs.OK() {
		rs.Decode(&ak)
	}

	if ak.AccessKey != "" && ak.AccessKey == id {
		set.Kind = "AccessKey"
		set.AccessKey = &ak
	} else {
		set.Error = types.NewErrorMeta(iamapi.ErrCodeNotFound, "Access Key Not Found")
	}
}

func (c AccessKey) ListAction() {

	ls := types.ObjectList{}
	defer c.RenderJson(&ls)

	k1 := iamapi.ObjKeyAccessKey(c.us.UserName, "zzzzzzzz")
	k2 := iamapi.ObjKeyAccessKey(c.us.UserName, "")
	if rs := store.Data.NewReader(nil).KeyRangeSet(k1, k2).
		ModeRevRangeSet(true).LimitNumSet(int64(ak_limit)).Query(); rs.OK() {

		for _, v := range rs.Items {
			var ak iamapi.AccessKey
			if err := v.Decode(&ak); err == nil {
				ls.Items = append(ls.Items, ak)
			}
		}
	}

	ls.Kind = "AccessKeyList"
}

func (c AccessKey) SetAction() {

	var set struct {
		types.TypeMeta
		iamapi.AccessKey
	}
	defer c.RenderJson(&set)

	if err := c.Request.JsonDecode(&set.AccessKey); err != nil {
		set.Error = types.NewErrorMeta(iamapi.ErrCodeInvalidArgument, "Bad Request")
		return
	}

	var prev iamapi.AccessKey
	if len(set.AccessKey.AccessKey) < 16 {
		set.AccessKey.AccessKey = iox_utils.Uint32ToHexString(uint32(time.Now().Unix())) + idhash.RandHexString(8)
	} else {

		if rs := store.Data.NewReader(
			iamapi.ObjKeyAccessKey(c.us.UserName, set.AccessKey.AccessKey)).Query(); rs.OK() {
			rs.Decode(&prev)
		}
	}

	if rs := store.Data.NewReader(nil).KeyRangeSet(
		iamapi.ObjKeyAccessKey(c.us.UserName, ""), iamapi.ObjKeyAccessKey(c.us.UserName, "")).
		LimitNumSet(int64(ak_limit + 1)).Query(); rs.OK() {
		if len(rs.Items) > ak_limit && prev.AccessKey == "" {
			set.Error = types.NewErrorMeta(iamapi.ErrCodeInvalidArgument, fmt.Sprintf("Num Out Range (%d)", ak_limit))
			return
		}
	}

	if prev.AccessKey == "" {
		prev = set.AccessKey
		prev.Created = types.MetaTimeNow()
	} else {

		prev.Action = set.AccessKey.Action
		prev.Description = set.AccessKey.Description

		for _, v := range set.AccessKey.Bounds {
			types.IterObjectLookup(prev.Bounds, v.Name, func(idx int) {
				if idx == -1 {
					v.Created = types.MetaTimeNow()
					prev.Bounds = append(prev.Bounds, v)
				}
			})
		}
	}

	if len(prev.SecretKey) < 40 {
		prev.SecretKey = idhash.RandBase64String(40)
	}

	if len(prev.User) < 1 {
		prev.User = c.us.UserName
	}

	if rs := store.Data.NewWriter(iamapi.ObjKeyAccessKey(c.us.UserName, prev.AccessKey), prev).
		Commit(); rs.OK() {
		set.Kind = "AccessKey"
	} else {
		set.Error = types.NewErrorMeta(iamapi.ErrCodeInternalError, "IO Error")
	}
}

func (c AccessKey) DelAction() {

	var set types.TypeMeta
	defer c.RenderJson(&set)

	id := c.Params.Get("access_key")
	if id == "" {
		set.Error = types.NewErrorMeta(iamapi.ErrCodeNotFound, "Access Key Not Found")
		return
	}

	if rs := store.Data.NewWriter(iamapi.ObjKeyAccessKey(c.us.UserName, id), nil).
		ModeDeleteSet(true).Commit(); rs.OK() {
		set.Kind = "AccessKey"
	} else {
		set.Error = types.NewErrorMeta(iamapi.ErrCodeInternalError, "IO Error")
	}
}

func (c AccessKey) BindAction() {

	var set types.TypeMeta
	defer c.RenderJson(&set)

	var (
		id    = c.Params.Get("access_key")
		bname = c.Params.Get("bound_name")
	)
	if id == "" && bname == "" {
		set.Error = types.NewErrorMeta(iamapi.ErrCodeNotFound, "Access Key Not Found")
		return
	}

	var ak iamapi.AccessKey
	if rs := store.Data.NewReader(iamapi.ObjKeyAccessKey(c.us.UserName, id)).Query(); rs.OK() {
		rs.Decode(&ak)
	}

	if id != ak.AccessKey {
		set.Error = types.NewErrorMeta(iamapi.ErrCodeNotFound, "Access Key Not Found")
		return
	}

	types.IterObjectLookup(ak.Bounds, bname, func(idx int) {
		if idx == -1 {
			ak.Bounds = append(ak.Bounds, iamapi.AccessKeyBound{
				Name:    bname,
				Created: types.MetaTimeNow(),
			})
		}
	})

	if rs := store.Data.NewWriter(iamapi.ObjKeyAccessKey(c.us.UserName, ak.AccessKey), ak).
		Commit(); rs.OK() {
		set.Kind = "AccessKey"
	} else {
		set.Error = types.NewErrorMeta(iamapi.ErrCodeInternalError, "IO Error")
	}
}

func (c AccessKey) UnbindAction() {

	var set types.TypeMeta
	defer c.RenderJson(&set)

	var (
		id    = c.Params.Get("access_key")
		bname = c.Params.Get("bound_name")
	)
	if id == "" && bname == "" {
		set.Error = types.NewErrorMeta(iamapi.ErrCodeNotFound, "Access Key Not Found")
		return
	}

	var ak iamapi.AccessKey
	if rs := store.Data.NewReader(iamapi.ObjKeyAccessKey(c.us.UserName, id)).Query(); rs.OK() {
		rs.Decode(&ak)
	}

	if id != ak.AccessKey {
		set.Error = types.NewErrorMeta(iamapi.ErrCodeNotFound, "Access Key Not Found")
		return
	}

	types.IterObjectLookup(ak.Bounds, bname, func(idx int) {
		if idx >= 0 {
			ak.Bounds = append(ak.Bounds[:idx], ak.Bounds[idx+1:]...)
		}
	})

	if rs := store.Data.NewWriter(iamapi.ObjKeyAccessKey(c.us.UserName, ak.AccessKey), ak).
		Commit(); rs.OK() {
		set.Kind = "AccessKey"
	} else {
		set.Error = types.NewErrorMeta(iamapi.ErrCodeInternalError, "IO Error")
	}
}
