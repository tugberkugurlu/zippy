using Microsoft.AspNet.Builder;
using Microsoft.Extensions.DependencyInjection;

namespace Zippy.Server 
{
    public class Startup 
    {
        public void ConfigureServices(IServiceCollection services) 
        {
            services.AddMvc();
        }
        
        public void Configure(IApplicationBuilder app)
        {
            app.UseMvc();
        }
    }
}