-module(parte2).
-export([
        start/0
        ]).

parse_input() ->
    {ok, Binary} = file:read_file("input.txt"),
    File = binary:bin_to_list(Binary),
    Passports = string:split(File, "\n\n", all),
    [[string:tokens(B, ":") || B <- string:tokens(A, "\n ")]
     || A <- Passports].

validate("byr", Field) -> 
    {Number, _} = string:to_integer(Field),
    (Number >= 1920) and (Number =< 2002);
validate("iyr", Field) -> 
    {Number, _} = string:to_integer(Field),
    (Number >= 2010) and (Number =< 2020);
validate("eyr", Field) -> 
    {Number, _} = string:to_integer(Field),
    (Number >= 2020) and (Number =< 2030);

validate("hgt", Field) -> 
    {Parsed, Metric} = string:to_integer(Field),
    case Metric of
        "cm" when Parsed >= 150, Parsed =< 193  -> true;
        "in" when Parsed >= 59, Parsed =< 76    -> true;
        _ -> false
    end;
validate("ecl", "amb")          -> true;
validate("ecl", "blu")          -> true;
validate("ecl", "brn")          -> true;
validate("ecl", "gry")          -> true;
validate("ecl", "grn")          -> true;
validate("ecl", "hzl")          -> true;
validate("ecl", "oth")          -> true;
validate("hcl", [$#|Hex] = _)   -> binary:decode_hex(binary:list_to_bin(Hex)), true;
validate("pid", Field)          -> string:to_integer(Field), length(Field) == 9;
validate("cid", _)              -> true;
validate(_, _)                  -> false.

validate_fields(Passport) -> 
    try 
        lists:all(fun([Field, Info]) -> validate(Field, Info) end, Passport)
    catch 
        _ -> false 
    end.

start() ->
    Passports = parse_input(),
    RequiredFields = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", 
                     "pid"],
    All_required_fields = fun(Passport) -> 
                                  Fields = [Field || [Field|_] <- Passport],
                                  lists:subtract(RequiredFields, Fields) =:= []
                          end,

    ValidPassports = [Passport || 
                      Passport <- lists:filter(All_required_fields, Passports), 
                      validate_fields(Passport)],
    erlang:display(length(ValidPassports)).
