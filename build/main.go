/*
 * Copyright 2018-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"os"

	"github.com/buildpack/libbuildpack/buildplan"
	"github.com/cloudfoundry/jvm-application-cnb/mainclass"
	"github.com/cloudfoundry/libcfbuildpack/build"
)

func main() {
	build, err := build.DefaultBuild()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to initialize Detect: %s\n", err)
		os.Exit(101)
	}

	if code, err := b(build); err != nil {
		build.Logger.Info(err.Error())
		os.Exit(code)
	} else {
		os.Exit(code)
	}
}

func b(build build.Build) (int, error) {
	build.Logger.FirstLine(build.Logger.PrettyIdentity(build.Buildpack))

	if m, ok, err := mainclass.NewMainClass(build); err != nil {
		return build.Failure(102), err
	} else if ok {
		if err = m.Contribute(); err != nil {
			return build.Failure(103), err
		}
	}

	return build.Success(buildplan.BuildPlan{})
}
