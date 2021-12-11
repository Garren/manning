defmodule Calculator do
  # define a default for b
  def sum(a, b \\ 0) do
    a + b
  end

  # this generates 3 functions
  # fun/3, fun/2, and fun/4
  def fun(a, b \\ 1, c, d \\ 2) do
    a + b + c + d
  end
end
