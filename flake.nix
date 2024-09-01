{
  description = "Ambiente de desenvolvimento em Go ";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        devShells.default = pkgs.mkShell
          {
            buildInputs = with pkgs; [
              go
              gotools
              gopls
              go-outline
              gopkgs
              godef
              golint
            ];

            shellHook = ''
              export GOPATH=$HOME/go
              export PATH=$PATH:$GOPATH/bin
            '';
          };
      }
    );
}
