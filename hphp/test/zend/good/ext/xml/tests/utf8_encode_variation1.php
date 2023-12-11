<?hh

class aClass {
   function __toString() :mixed{
       return "Some Ascii Data";
   }
}
<<__EntryPoint>>
function entrypoint_utf8_encode_variation1(): void {
  /* Prototype  : proto string utf8_encode(string data)
   * Description: Encodes an ISO-8859-1 string to UTF-8
   * Source code: ext/xml/xml.c
   * Alias to functions:
   */

  echo "*** Testing utf8_encode() : usage variations ***\n";
  error_reporting(E_ALL & ~E_NOTICE);

  // Initialise function arguments not being substituted (if any)


  //array of values to iterate over
  $values = vec[

        // int data
        0,
        1,
        12345,
        -2345,

        // float data
        10.5,
        -10.5,
        10.1234567e10,
        10.7654321E-10,
        .5,








        // null data
        NULL,
        null,

        // boolean data
        true,
        false,
        TRUE,
        FALSE,

        // empty data
        "",
        '',

        // object data
        new aClass(),


  ];

  // loop through each element of the array for data

  foreach($values as $value) {
        echo @"\nArg value $value \n";
        try { var_dump( utf8_encode($value) ); } catch (Exception $e) { echo "\n".'Warning: '.$e->getMessage().' in '.__FILE__.' on line '.__LINE__."\n"; }
  }

  echo "Done";
}
