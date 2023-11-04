-module(parte1).
-export([start/0]).

% usei hashmaps pra montar a estrutura dos diretórios porque não sabia (ainda não sei) 
% como fazer uma arvore em erlang.
% o caminho absoluto dos diretórios servem como chave e o valor é o conteudo deles.
% tentei fazer com que o nome do diretório fossem as chaves porem acabei descobrindo
% que alguns diretórios tem nomes repetidos.

parse_input() ->
    {ok, Binary} = file:read_file("dirs.txt"),
    File = string:trim(binary:bin_to_list(Binary)),
    Lines = string:split(File, "\n", all),
    [string:split(X, " ", all) || X <- Lines].

read_ls([], Files) ->
    {Files, []};
read_ls([["$" | _] | _] = Commands, Files) ->
    {Files, Commands};
read_ls([File | T], Files) ->
    read_ls(T, [File | Files]).
read_ls(T) ->
    read_ls(T, []).


build_tree([], _, Tree) -> Tree;

build_tree([["$", "cd", ".."] | T], [_|Parent], Tree) ->
    build_tree(T, Parent, Tree);

build_tree([["$", "cd", Dir] | T], Pwd, Tree) ->
    build_tree(T, [Dir | Pwd], Tree);

build_tree([["$", "ls"] | T], Pwd, Tree) ->
    {Files, Rest} = read_ls(T),
    build_tree(Rest, Pwd, Tree#{Pwd => Files}).

build_tree([_|T]) ->
    build_tree(T, ["/"], #{}).

count_sizes(Tree, Pwd) ->
    #{Pwd := Files} = Tree,
    lists:sum([ 
     case X of
         ["dir", Dir] -> count_sizes(Tree, [Dir | Pwd]);
         [Size, _] -> 
             {Int, _} = string:to_integer(Size), Int
     end
     || X <- Files]).

start() ->
    Input = parse_input(),
    Tree = build_tree(Input),
    count_sizes(Tree, ["/"]),
    % isso recalcula o tamanho de cada diretório a cada iteração
    % ineficiente mas da pro gasto.
    Sizes = [count_sizes(Tree, X) || X <- maps:keys(Tree)],
    erlang:display(lists:sum([X || X <- Sizes, X < 100000])).

