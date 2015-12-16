using System;
using System.IO;
using System.Threading.Tasks;
using Zippy.Server.Abstractions;

namespace Zippy.Server.FileSystem
{
    public class FileSystemPackageReadRepository : IPackageReadRepository 
    {
        private readonly IFileSystem _fileSystem;
        
        public FileSystemPackageReadRepository(IFileSystem fileSystem) 
        {
            if (fileSystem == null)
            {
                throw new ArgumentNullException(nameof(fileSystem));   
            }
            
            _fileSystem = fileSystem;
        }

        public Task<IPackage> Exists(PackageName name)
        {
            throw new NotImplementedException();
        }

        public Task<IPackage> Exists(PackageName name, SemanticVersion version)
        {
            throw new NotImplementedException();
        }

        public Task<PaginatedResult<IPackage>> Get(long skip, byte take)
        {
            throw new NotImplementedException();
        }

        public Task<IPackage> Get(PackageName name, SemanticVersion version)
        {
            throw new NotImplementedException();
        }

        public Task<Stream> GetStream(PackageName name, SemanticVersion version)
        {
            throw new NotImplementedException();
        }
    }
}