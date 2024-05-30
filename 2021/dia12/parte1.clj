(ns huehue
  (:require [clojure.string :as str]))

(->> (str/split-lines (slurp "input.txt"))
     (map (fn [x] (str/split x #"-")))
     (flatten)
     (apply hash-map)
     (print)
     )

;; (print paths)
