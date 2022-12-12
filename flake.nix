{
  description = "Advent of Code solutions";

  inputs.nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
      in {
        devShell = with pkgs; mkShell rec {
          name = "adventofcode";
          buildInputs = [
            delve
            graphviz
            go
            gopls
          ];
        };
      });
}
