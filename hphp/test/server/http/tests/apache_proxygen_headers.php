<?hh

<<__EntryPoint>> function main(): void {
  require_once('test_base.inc');
  init();
  requestAll(
    vec[
      vec[
        'test_proxygen_headers.php',
        null,
        dict[
          'xyzzy' => 42,
          'XyZZy' => 43,
          'XYZZY' => 44,
          'xxxxx' => 45,
        ]
      ]
    ],
  );
}
