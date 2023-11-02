// MIT License
//
// Copyright (c) 2018 SpiralScout
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package php

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

// Generate generates needed service classes
func Generate(req *plugin.CodeGeneratorRequest) *plugin.CodeGeneratorResponse {
	resp := &plugin.CodeGeneratorResponse{}
	resp.SupportedFeatures = proto.Uint64(uint64(plugin.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL))

	for _, file := range req.ProtoFile {
		for _, service := range file.Service {
			resp.File = append(resp.File, generate(req, file, service))
		}
	}

	for _, file := range req.ProtoFile {
		for _, service := range file.Service {
			resp.File = append(resp.File, generate1(req, file, service))
		}
	}

	return resp
}

func generate(
	req *plugin.CodeGeneratorRequest,
	file *descriptor.FileDescriptorProto,
	service *descriptor.ServiceDescriptorProto,
) *plugin.CodeGeneratorResponse_File {
	return &plugin.CodeGeneratorResponse_File{
		Name:    str(filename(file, service.Name)),
		Content: str(body(req, file, service)),
	}
}

func generate1(
	req *plugin.CodeGeneratorRequest,
	file *descriptor.FileDescriptorProto,
	service *descriptor.ServiceDescriptorProto,
) *plugin.CodeGeneratorResponse_File {
	return &plugin.CodeGeneratorResponse_File{
		Name:    str(filename1(file, service.Name)),
		Content: str(body1(req, file, service)),
	}
}

// helper to convert string into string pointer
func str(str string) *string {
	return &str
}
