{
  sources ? import ./npins,
  system ? builtins.currentSystem,
  pkgs ? import sources.nixpkgs {
    inherit system;
    config = { };
    overlays = [ ];
  },
}:

pkgs.mkShell {
  packages = with pkgs; [
    go
    just
    go-tools
    golangci-lint
  ];
}
