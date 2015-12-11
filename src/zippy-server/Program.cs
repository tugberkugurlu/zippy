using System;
using Microsoft.AspNet.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.PlatformAbstractions;

namespace Zippy.Server 
{
    public class Program
    {
        public static void Main(string[] args)
        {   
            var appEnv = PlatformServices.Default.Application;
            var runtimeEnv = PlatformServices.Default.Runtime;
            
            var config = new ConfigurationBuilder()
                .AddJsonFile("hosting.json")
                .AddCommandLine(args)
                .Build();
                
            using (new WebHostBuilder(config).Build().Start())
            {
                Console.WriteLine("Started the server..");
                Console.WriteLine("Press any key to stop the server");
                Console.ReadLine();
            }
        }
    }    
}