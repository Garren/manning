defmodule TestPrivate do
  # functions are exported by default
  def double(a), do: sum(a, a)

  # defp specifies a private function
  defp sum(a, b), do: a + b
end
