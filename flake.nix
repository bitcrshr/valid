{
  description = "dev env for valid module :)";

  inputs = {
    # go 1.24.4
    nixpkgs.url = "github:nixos/nixpkgs?ref=08f22084e6085d19bcfb4be30d1ca76ecb96fe54";
    flake-utils = {
      url = "github:numtide/flake-utils";
    };
  };

  outputs = { self, nixpkgs, flake-utils }:
  flake-utils.lib.eachDefaultSystem (system:
    let
      pkgs = import nixpkgs {
        inherit system;
      };
    in
    with pkgs;
    {
      devShells.default = mkShell {
        buildInputs = [
          gh
          go
          gopls
          gofumpt
          golangci-lint
          golangci-lint-langserver
        ];

        shellHook = ''
        '';
      };
    }
  );

}
