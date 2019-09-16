/*
 * Copyright 2019 The Baudtime Authors
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package visitor

import (
	"github.com/baudtime/baudtime/meta"
	"github.com/baudtime/baudtime/msg"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

func init() {
	registry["slavesonly"] = slavesOnly
}

func slavesOnly(shard *meta.Shard, op func(node *meta.Node) (resp msg.Message, err error)) (resp msg.Message, err error) {
	if len(shard.Slaves) == 0 {
		return nil, errors.Errorf("shard %v has no available slaves", shard.ID)
	}

	var multiErr error

	for _, node := range shard.Slaves {
		if resp, err = op(node); err != nil {
			multiErr = multierror.Append(multiErr, err)
		} else {
			return
		}
	}

	return nil, multiErr
}
