<?hh

// This entire file is formatted because it does not contain the
// "partially-generated" tag.

class Foo {
  function a () : int { return 0; }
  /* BEGIN MANUAL SECTION */
  function b () : int { return 0; }
  function c () : int { return 0; }
  /* END MANUAL SECTION */
  function d () : int { return 0; }

  /* BEGIN MANUAL SECTION */
  function e () : int {
    /* HH_FIXME[4110] */ return vec[ 1 ,
                                     2 ,
                                     3 ];
  }
  /* END MANUAL SECTION */
}
