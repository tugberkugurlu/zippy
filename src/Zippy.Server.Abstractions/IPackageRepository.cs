using System;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;

namespace Zippy.Server.Abstractions
{
    public class PaginatedResult<TItem>
    {
        public PaginatedResult(long skipped, byte taken, long totalCount, IReadOnlyList<TItem> items) 
        {
            if(items == null)
            {
                throw new ArgumentNullException(nameof(items));
            }
            
            Skipped = skipped;
            Taken = taken;
            TotalCount = totalCount;
            Items = items;
        }
        
        public long Skipped { get; }
        public byte Taken { get; }
        public long TotalCount { get; }
        public IReadOnlyList<TItem> Items { get; }
    }
    
    public interface IPackageRepository
    {
        Task<IPackage> Get(PackageName name, SemanticVersion version);
        Task<IPackage> Exists(PackageName name, SemanticVersion version);
        Task<IPackage> Exists(PackageName name);
        Task<PaginatedResult<IPackage>> Get(long skip, byte take);
        Task<Stream> GetStream(PackageName name, SemanticVersion version);
    }
    
    public interface IPackageSearchManager 
    {
        Task<PaginatedResult<IPackage>> Search(string lookupTerm, long skip, byte take);
    }
    
    public static class PackageExtensions 
    {
    }
}