<?hh
<<__EntryPoint>> function main(): void {
$tests = vec[
    "\x41\xC2\x3E\x42",
    "\xE3\x80\x22",
    "\x41\x98\xBA\x42\xE2\x98\x43\xE2\x98\xBA\xE2\x98",
];
foreach ($tests as $t) {
    echo bin2hex(utf8_decode($t)), "\n";
}
echo "Done.\n";
}
