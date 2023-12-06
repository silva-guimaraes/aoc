-module(parte2).
-export([start/0]).

% https://adventofcode.com/2023/day/5

% esse foi o problema mais dificil até agora

% em minhas experiencias com essa linguagem não houve nenhuma vez onde eu me
% surpreendi com a velocidade da lingua, muito pelo contrario eu ja esperava
% sempre que a solução seria lerda

% eu tinha uma ideia em mente de como resolver essa segunda parte com força bruta, 
% mas eu também acreditava que resolver do jeito que eu fiz (do jeito esperto) seria facíl. 
% não foi. só consegui a resposta correta faltando uma hora pro proximo desafio. 
% isso não teria sido nada se eu não tivesse passado o dia inteiro quebrando a 
% cabeça com isso daqui. logo no dia seguinte, logo na segunda parte 
% do problema eu tive outra oportunidade de escolher entre usar força bruta ou pensar 
% em algum algoritmo inteligente. não medi esforços e usei força bruta. 
% fiquei traumatizado com essa segunda parte aqui de hoje.
% não é a primeira vez que isso acontece e nem sera a ultima.
% essa minha fabula é um conto pra nunca se esquecer.

% -----


pair([]) -> [];
pair([Start, Range | T]) ->
    % juntar inicio com tamanho para forma o fim da faixa
    % facilita a vida
    [{Start, Start + Range} | pair(T)].

parse_input() ->
    {ok, Binary} = file:read_file("input.txt"),
    [_Seeds | Sections] = 
    string:split( string:trim( binary:bin_to_list(Binary)), "\n\n", all),
    Maps =
    [ begin 
          [_ | B] = string:split(X, ":\n"), 
          C = string:split(B, "\n", all),
          [
           begin
               E = string:tokens(D, " "),
               lists:map(fun(F) -> {G, _} = string:to_integer(F), G end, E)
           end
           || D <- C 
          ]
      end 
      || X <- Sections],
    Seeds = 
    begin
        [_ | [A1]] = string:split(_Seeds, ": "),
        B1 = string:tokens(A1, " "),
        C1 = [ begin {Y, _} = string:to_integer(X), Y end || X <- B1],
        pair(C1)
    end,
    {Seeds, Maps}.


pass(Seed, [Dest, Source, _]) ->
    Dest + (Seed - Source).

% iterar por todos os maps e encontrar alguma que se encaixe com a faixa.
% as faixas ja foram cortadas então não acontece de uma faixa cair em mais de um map

% não alterar valor caso nenhuma map que se encaixe for encontrado
find_range(Range, []) -> Range;

find_range({Start, End}, [[_, Source, Range] = Line | _]) 
  when (Start >= Source) and (Start < Source + Range) and
       (End >= Source) and (End < Source + Range) ->
    [{pass(Start, Line), pass(End, Line)}];

find_range(Range, [_ | Lines]) ->
    find_range(Range, Lines).

% essas funç̃oes separam uma faixas eu varias outras novas caso elas caiam
% não 100% dentro de um map
slice_range(Range, []) -> [Range];

% faixa se inicia dentro e termina depois de um map
slice_range({Start, End}, [[_, Source, Range] | _]) 
  when (Source =< Start) and (Start < Source + Range) and
       (End >= Source + Range) ->
    [
     {Start, Source + Range - 1}, 
     {Source + Range, End}
    ];

% faixa se inicia antes e termina dentro de um map
slice_range({Start, End}, [[_, Source, Range] | _]) 
  when (Start < Source) and 
       (End >= Source) and (End < Source + Range) ->
    [
     {Start, Source - 1}, 
     {Source, End}
    ];

% faixa totalmente dentro de um map significa nenhuma outra faixa nova
% apenas manter a faixa do jeito que esta
slice_range({Start, End}, [[_, Source, Range] | _]) 
  when (Start >= Source) and (Start < Source + Range) and
       (End >= Source) and (End < Source + Range) ->
    [{Start, End}];

% faixa cobre completamente um map.
% isso separa a faixa em 3 faixas novas
slice_range({Start, End}, [[_, Source, Range] | _]) 
  when (Start < Source) and (End > Source + Range) ->
    [
     {Source, Source + Range - 1}, 
     {Start, Source - 1},
     {Source + Range, End}
    ];

slice_range(Range, [_ | Lines]) ->
    slice_range(Range, Lines).

loop(Seeds, []) -> Seeds;
loop(Seeds, [Map | Maps] = M) ->
    A = lists:map(fun(X) -> slice_range(X, Map) end, Seeds),
    B = lists:flatten(A),
    % demorei um tempo pra perceber que depois de recortar as faixas essas
    % faixas novas poderiam cair em algum outro map diferente
    case lists:flatten(lists:map(fun(X) -> slice_range(X, Map) end, B)) of 
        X when length(B) =/= length(X) ->
            % repetir loop com novas faixas quando isso acontecer
            loop(X, M);
        _ -> 
            C = lists:map(fun(X) -> find_range(X, Map) end, B),
            loop(lists:flatten(C), Maps)
    end.

start() ->
    {Seeds, Maps } = parse_input(),
    Foo = loop(Seeds, Maps),
    erlang:display(lists:min([ X || {X, _} <- Foo])).
