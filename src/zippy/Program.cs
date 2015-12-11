using System.Reflection;
using Microsoft.Dnx.Runtime.Common.CommandLine;

namespace Zippy
{
    public class Program
    {
        public int Main(string[] args)
        {
            var app = new CommandLineApplication();
            app.Name = "zippy";
            app.FullName = "zippy - Distribute Your Stuff";
            app.Description = "General purpose package distribution tool for your stuff (yes, all your stuff™)";
            app.HelpOption("-?|-h|--help");
            app.VersionOption("-v|--version", () => GetVersion());

            // Show help information if no subcommand/option was specified
            app.OnExecute(() =>
            {
                app.ShowHelp();
                return 2;
            });

            RegisterPackCommand(app);
            RegisterPushCommand(app);
            RegisterInstallCommand(app);

            return app.Execute(args);
        }

        private static void RegisterPackCommand(CommandLineApplication cmdApp)
        {
            cmdApp.Command("pack", (c) =>
            {
                c.Description = "Creates the package based on the given manifest file.";
                c.HelpOption("-?|-h|--help");

                c.OnExecute(() =>
                {
                    return -1;
                });
            });
        }

        private static void RegisterPushCommand(CommandLineApplication cmdApp)
        {
            cmdApp.Command("push", (c) =>
            {
                c.Description = "Pushes the zippy package to given zippy endpoint.";
                c.HelpOption("-?|-h|--help");

                c.OnExecute(() =>
                {
                    return -1;
                });
            });
        }

        private static void RegisterInstallCommand(CommandLineApplication cmdApp)
        {
            cmdApp.Command("install", (c) =>
            {
                c.Description = "Installs the zippy package in given directory from the given source.";
                c.HelpOption("-?|-h|--help");

                c.Argument("[package]", "The name of the package.");

                var versionOption = c.Option("-v|--version <VERSION>", "The version of the package to install. The version number needs to be a valid SemVer.", CommandOptionType.SingleValue);
                var sourceOption = c.Option("-s|--source <SOURCE>", "A list of sources to use for zippy endpoints to be used to install the package.", CommandOptionType.MultipleValue);
                var outputDirectoryOption = c.Option("-out|--outputDirectory <OUTPUT-DIRECTORY>", "The output directory to extract the package.", CommandOptionType.SingleValue);

                c.OnExecute(() =>
                {
                    return -1;
                });
            });
        }

        private static string GetVersion()
        {
            var assembly = typeof(Program).GetTypeInfo().Assembly;
            var assemblyInformationalVersionAttribute = assembly.GetCustomAttribute<AssemblyInformationalVersionAttribute>();
            return assemblyInformationalVersionAttribute.InformationalVersion;
        }
    }
}
