// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package traceidextractor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/traceidextractor"

import (
	"context"

	"go.opentelemetry.io/collector/model/pdata"
)

type traceidprocessor struct {}

// newTracesProcessor returns a processor that modifies attributes of a span.
// To construct the attributes processors, the use of the factory methods are required
// in order to validate the inputs.
func newTraceIDExtractor() *traceidprocessor {
	return &traceidprocessor{}
}

func (a *traceidprocessor) processTraces(_ context.Context, td pdata.Traces) (pdata.Traces, error) {
	rss := td.ResourceSpans()
	for i := 0; i < rss.Len(); i++ {
		rs := rss.At(i)
		ilss := rs.InstrumentationLibrarySpans()
		for j := 0; j < ilss.Len(); j++ {
			ils := ilss.At(j)
			spans := ils.Spans()
			for k := 0; k < spans.Len(); k++ {
				span := spans.At(k)
				span.Attributes().Insert("otlp_trace_id", pdata.NewAttributeValueString(span.TraceID().HexString()))
				}
		}
	}
	return td, nil
}