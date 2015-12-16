namespace Zippy.Server.FileSystem
{
    public interface IFileSystem 
    {
        bool FileExists(string path);
    }
}