// +build !ignore_autogenerated

/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package wholepkg

import (
	conversion "k8s.io/kubernetes/pkg/conversion"
	reflect "reflect"
)

// GetGeneratedDeepCopyFuncs returns the generated funcs, since we aren't registering them.
func GetGeneratedDeepCopyFuncs() []conversion.GeneratedDeepCopyFunc {
	return []conversion.GeneratedDeepCopyFunc{
		{Fn: DeepCopy_wholepkg_Struct_B, InType: reflect.TypeOf(&Struct_B{})},
		{Fn: DeepCopy_wholepkg_Struct_Empty, InType: reflect.TypeOf(&Struct_Empty{})},
		{Fn: DeepCopy_wholepkg_Struct_Primitives, InType: reflect.TypeOf(&Struct_Primitives{})},
		{Fn: DeepCopy_wholepkg_Struct_Primitives_Alias, InType: reflect.TypeOf(&Struct_Primitives_Alias{})},
		{Fn: DeepCopy_wholepkg_Struct_Primitives_Embed, InType: reflect.TypeOf(&Struct_Primitives_Embed{})},
	}
}

func DeepCopy_wholepkg_Struct_B(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Struct_B)
		out := out.(*Struct_B)
		_ = in
		_ = out
		return nil
	}
}

func DeepCopy_wholepkg_Struct_Empty(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Struct_Empty)
		out := out.(*Struct_Empty)
		_ = in
		_ = out
		return nil
	}
}

func DeepCopy_wholepkg_Struct_Primitives(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Struct_Primitives)
		out := out.(*Struct_Primitives)
		out.BoolField = in.BoolField
		out.IntField = in.IntField
		out.StringField = in.StringField
		out.FloatField = in.FloatField
		return nil
	}
}

func DeepCopy_wholepkg_Struct_Primitives_Alias(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Struct_Primitives_Alias)
		out := out.(*Struct_Primitives_Alias)
		out.BoolField = in.BoolField
		out.IntField = in.IntField
		out.StringField = in.StringField
		out.FloatField = in.FloatField
		return nil
	}
}

func DeepCopy_wholepkg_Struct_Primitives_Embed(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Struct_Primitives_Embed)
		out := out.(*Struct_Primitives_Embed)
		out.Struct_Primitives = in.Struct_Primitives
		return nil
	}
}
