-module(parte1).
-export([
        start/0,
        parse_input/0
        ]).

parse_input() ->
    {ok, Binary} = file:read_file("input.txt"),
    File = binary:bin_to_list(Binary),
    Passports = string:split(File, "\n\n", all),
    [[begin [C, _] = string:tokens(B, ":"), C end
       || B <- string:tokens(A, "\n ")]
     || A <- Passports].


start() ->
    Passports = parse_input(),
    RequiredFields = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", 
                     "pid"], %% "cid"
    Valid = [X || X <- Passports, 
          lists:subtract(RequiredFields, X) =:= []
    ],
    erlang:display(length(Valid)).


    % Path = path(Forest, 0, 0),
    % io:format("~p~n", [length([X || X <- Path, X == $#])]).

