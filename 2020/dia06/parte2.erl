-module(parte2).
-export([
        start/0
        ]).

parse_input() ->
    {ok, Binary} = file:read_file("input.txt"),
    File = string:trim(binary:bin_to_list(Binary)),
    Groups = string:split(File, "\n\n", all),
    [string:split(X, "\n", all) || X <- Groups].



count_person_answers([], Count) -> Count;

count_person_answers([YesQuestion|T], Count) ->
    New = #{YesQuestion => 1},
    Merged = maps:merge_with(fun(_, A, B) -> A+B end, New, Count),
    count_person_answers(T, Merged).

count_person_answers(YesQuestions) -> count_person_answers(YesQuestions, #{}).



count_group_answers([], GroupSize, Count) ->
    length([K || K := V <- Count, V == GroupSize]);

count_group_answers([Person|T], GroupSize, Count) -> 
    New = count_person_answers(Person),
    Merged = maps:merge_with(fun(_, A, B) -> A+B end, New, Count),
    count_group_answers(T, GroupSize, Merged).

count_group_answers(People) ->
    count_group_answers(People, length(People), #{}).



start() ->
    Groups = parse_input(),
    erlang:display(lists:sum([count_group_answers(Group) || Group <- Groups])).



