// Code generated by github.com/hooto/protobuf_slice
// source: types.proto
// DO NOT EDIT!

package iamapi

import "sync"

var object_slice_mu_MsgItem sync.RWMutex

func (it *MsgItem) Equal(it2 *MsgItem) bool {
	if it2 == nil ||
		it.Id != it2.Id ||
		it.Action != it2.Action ||
		it.ToUser != it2.ToUser ||
		it.ToEmail != it2.ToEmail ||
		it.FromUser != it2.FromUser ||
		it.FromEmail != it2.FromEmail ||
		it.Priority != it2.Priority ||
		it.Title != it2.Title ||
		it.Body != it2.Body ||
		it.Retry != it2.Retry ||
		it.Log != it2.Log ||
		it.Posted != it2.Posted ||
		it.Created != it2.Created ||
		it.Updated != it2.Updated {
		return false
	}
	return true
}

func (it *MsgItem) Sync(it2 *MsgItem) bool {
	if it2 == nil {
		return false
	}
	if it.Equal(it2) {
		return false
	}
	*it = *it2
	return true
}

func MsgItemSliceGet(ls []*MsgItem, arg_id string) *MsgItem {
	object_slice_mu_MsgItem.RLock()
	defer object_slice_mu_MsgItem.RUnlock()

	for _, v := range ls {
		if v.Id == arg_id {
			return v
		}
	}
	return nil
}

func MsgItemSliceDel(ls []*MsgItem, arg_id string) ([]*MsgItem, bool) {
	object_slice_mu_MsgItem.Lock()
	defer object_slice_mu_MsgItem.Unlock()
	for i, v := range ls {
		if v.Id == arg_id {
			ls = append(ls[:i], ls[i+1:]...)
			return ls, true
		}
	}
	return ls, false
}

func MsgItemSliceEqual(ls, ls2 []*MsgItem) bool {
	object_slice_mu_MsgItem.RLock()
	defer object_slice_mu_MsgItem.RUnlock()

	if len(ls) != len(ls2) {
		return false
	}
	hit := false
	for _, v := range ls {
		hit = false
		for _, v2 := range ls2 {
			if v.Id != v2.Id {
				continue
			}
			if !v.Equal(v2) {
				return false
			}
			hit = true
			break
		}
		if !hit {
			return false
		}
	}
	return true
}

func MsgItemSliceSync(ls []*MsgItem, it2 *MsgItem) ([]*MsgItem, bool) {
	if it2 == nil {
		return ls, false
	}
	object_slice_mu_MsgItem.Lock()
	defer object_slice_mu_MsgItem.Unlock()

	hit := false
	changed := false
	for i, v := range ls {
		if v.Id != it2.Id {
			continue
		}
		if !v.Equal(it2) {
			ls[i], changed = it2, true
		}
		hit = true
		break
	}
	if !hit {
		ls = append(ls, it2)
		changed = true
	}
	return ls, changed
}

func MsgItemSliceSyncSlice(ls, ls2 []*MsgItem) ([]*MsgItem, bool) {
	if MsgItemSliceEqual(ls, ls2) {
		return ls, false
	}
	return ls2, true
}