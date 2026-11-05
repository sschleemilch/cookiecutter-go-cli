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
    go_1_27
    just
    go-tools
    golangci-lint
  ];
}
