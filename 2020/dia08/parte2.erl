-module(parte2).
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

execute(Ops, Pos, Counter, _) when Pos > length(Ops) -> {ok, Counter};
execute(Ops, Pos, Counter, Visited) ->
    {Increment, Offset} = case lists:nth(Pos, Ops) of
                               {jmp, Jump}      -> {0, Jump};
                               {acc, Amount}    -> {Amount, 1};
                               {nop, _}         -> {0, 1}
                           end,
    case lists:member(Pos, Visited) of
        true -> {loop, Counter};
        false ->
           execute(Ops, Pos + Offset, Counter + Increment, [Pos|Visited])
    end.
execute(Ops) -> execute(Ops, 1, 0, []).

nop_jump([], _)                 -> [];
nop_jump([{jmp, _}|T], 1)       -> [{nop, 0} | nop_jump(T, 0)];
nop_jump([H|T],         Nth)    -> [H | nop_jump(T, Nth-1)].

is_jump({_, {jmp, _}})      -> true;
is_jump(_)                  -> false.

%% não funciona caso a instrução errada seja um nop.
%% acabei não percebendo isso e por sorte/azar o meu input 
%% tinha um jmp como a instrução incorreta.
bruteforce(Ops) -> 
    Enumerate = lists:enumerate(Ops),
    Jumps = lists:filter(fun is_jump/1, Enumerate),
    Results = [ execute(nop_jump(Ops, Jump)) || {Jump, _} <- Jumps ],
    [X || {ok, X} <- Results].

start() ->
    Ops = parse_input(),
    % Jumps = jump_indexes(Ops),
    [Ret|_] = bruteforce(Ops),
    erlang:display(Ret).



