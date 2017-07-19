//  Copyright 2017 Google Inc. All Rights Reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package workflow

import (
	"fmt"
	"regexp"
)

var (
	instances      = map[*Workflow]*resourceMap{}
	instanceURLRgx = regexp.MustCompile(fmt.Sprintf(`^(projects/(?P<project>%[1]s)/)?zones/(?P<zone>%[1]s)/instances/(?P<instance>%[1]s)$`, rfc1035))
)

func initInstancesMap(w *Workflow) {
	m := &resourceMap{}
	instances[w] = m
	m.typeName = "instance"
	m.urlRgx = instanceURLRgx
}

func deleteInstance(w *Workflow, r *resource) error {
	if err := w.ComputeClient.DeleteInstance(w.Project, w.Zone, r.real); err != nil {
		return err
	}
	r.deleted = true
	return nil
}