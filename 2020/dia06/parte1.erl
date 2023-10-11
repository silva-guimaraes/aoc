-module(parte1).
-export([
        start/0
        ]).

parse_input() ->
    {ok, Binary} = file:read_file("input.txt"),
    File = string:trim(binary:bin_to_list(Binary)),
    string:split(File, "\n\n", all).

count_group_answers(Group) -> 
    length(lists:delete($\n, lists:uniq(Group))).

start() ->
    Groups = parse_input(),
    erlang:display(lists:sum([count_group_answers(Group) || Group <- Groups])).



