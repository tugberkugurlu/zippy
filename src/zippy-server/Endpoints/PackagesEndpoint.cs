using Microsoft.AspNet.Mvc;

namespace Zippy.Server.Endpoints 
{
    [Route("packages")]
    public class PackagesEndpoint : Controller 
    {
        [HttpGet]
        public string[] Get()
        {
            return new[]
            {
                "Package 1",
                "Package 2"
            };
        }
    }
}