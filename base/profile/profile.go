// Copyright 2015 lessOS.com, All rights reserved.
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

package profile

import (
	"errors"
	"strings"
	"time"

	"../../idsapi"
)

func PutValidate(set idsapi.UserProfile) (idsapi.UserProfile, error) {

	set.Name = strings.TrimSpace(set.Name)
	if len(set.Name) < 1 || len(set.Name) > 30 {
		return set, errors.New("Name must be between 1 and 30 characters long")
	}

	if _, err := time.Parse("2006-01-02", set.Birthday); err != nil {
		return set, errors.New("Birthday is not valid")
	}

	if set.About == "" {
		return set, errors.New("About Me can not be null")
	}

	return set, nil
}
