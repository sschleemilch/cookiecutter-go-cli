let
  nixpkgs = fetchTarball "https://github.com/NixOS/nixpkgs/tarball/nixos-26.05";
  pkgs = import nixpkgs {
    config = { };
    overlays = [ ];
  };
in

pkgs.mkShell {
  packages = with pkgs; [
    go
    just
    go-tools
    golangci-lint
  ];

  # LD_LIBRARY_PATH = pkgs.lib.makeLibraryPath [ pkgs.stdenv.cc.cc.lib ];
}
