<?hh
// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.

<<__SupportDynamicType>>
class C { }
function returnSD(?C $c): C {
  if ($c is nonnull) {
    return $c;
  }
  else return new C();
}
