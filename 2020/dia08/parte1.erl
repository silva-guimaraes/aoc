-module(parte1).
-export([
        start/0
        ]).

parse_input() ->
    {ok, Binary} = file:read_file("input.txt"),
    File = string:trim(binary:bin_to_list(Binary)),
    Lines = string:split(File, "\n", all),
    [ begin
          [Op, StringNumber] = string:tokens(X, " "),
          {Number, []} = string:to_integer(StringNumber),
          {erlang:list_to_atom(Op), Number}
      end
      || X <- Lines].

execute(Ops, Pos, Counter, Visited) ->
    {Increment, Offset} = case lists:nth(Pos, Ops) of
                               {jmp, Jump}      -> {0, Jump};
                               {acc, Amount}    -> {Amount, 1};
                               {nop, _}         -> {0, 1}
                           end,
    case lists:member(Pos, Visited) of
        true -> Counter;
        false ->
           execute(Ops, Pos + Offset, Counter + Increment, [Pos|Visited])
    end.

execute(Ops) -> execute(Ops, 1, 0, []).


start() ->
    Input = parse_input(),
    execute(Input).



