<?hh <<__EntryPoint>> function main(): void {
$values = vec[10,
                3950.5,
                3.9505e3,
                03,
                0x5F,
                "10",
                "3950.5",
                "3.9505e3",
                "039",
                "0x5F",
                true,
                false,
                null,
                ];

for ($i = 0; $i < count($values); $i++) {
    $res = decoct($values[$i]);
    var_dump($res);
}
}
