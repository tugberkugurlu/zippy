using System;
using System.IO;

namespace Zippy.Server.Abstractions
{
    public class PackageName
    {
        public PackageName(string name)
        {
            if(name == null)
            {
                throw new ArgumentNullException(nameof(name));
            }
            
            if (!FileNameValidator.IsValid(name))
            {
                throw new InvalidOperationException($"TODO: Package Name is not a valid package name: {name}"));
            }

            Value = name;
        }

        public string Value { get; }

        private static class FileNameValidator
        {
            public static bool IsValid(string fileName)
            {
                return (fileName.IndexOfAny(Path.GetInvalidFileNameChars()) >= 0) == false;
            }
        }
    }
}