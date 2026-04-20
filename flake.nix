{
  description = "rme — a command launcher for RUNME.sh";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      systems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
      forAllSystems = nixpkgs.lib.genAttrs systems;
    in
    {
      packages = forAllSystems (system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        {
          default = pkgs.buildGoModule rec {
            pname = "rme";
            version = "0.1.0";

            src = ./.;

            vendorHash = null;

            ldflags = [
              "-s"
              "-w"
              "-X main.version=${version}"
            ];

            meta = with pkgs.lib; {
              description = "A command launcher for RUNME.sh with shell completions";
              homepage = "https://github.com/mipmip/rme";
              license = licenses.mit;
              maintainers = [ ];
              mainProgram = "rme";
            };
          };

          rme = self.packages.${system}.default;
        }
      );

      apps = forAllSystems (system: {
        default = {
          type = "app";
          program = "${self.packages.${system}.default}/bin/rme";
        };

        rme = self.apps.${system}.default;
      });

      devShells = forAllSystems (system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [
              go
              gopls
              gotools
              go-tools
            ];

            shellHook = ''
              echo "rme development environment"
              echo "Go version: $(go version)"
            '';
          };
        }
      );
    };
}
