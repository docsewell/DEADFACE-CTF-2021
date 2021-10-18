using System;
using System.Security.Cryptography;
using System.Text;

public static class SymmetricEncryptor
{
    // don't use this

    //static string password = "very strong password 123412;,[p;[; 172634812";
    static string password = "f1ag{you-didnt-think-it-was-that-easy-did-you}";
    
    public static void Main(string[] args)
    {
        Console.WriteLine("hlS4MbOmA+kQX71xXwPs7CsCWp9jQxCPa/oMk2o2bZr+jgweD4b8u80z5LVoBqC7"); 
    }
    
    public static byte[] EncryptString(string toEncrypt)
    {
        //var key = GetKey(password);
        byte[] key = {0x05,0x12,0x3D,0x2C,0x7D,0x22,0xF7,0x5A,0x9B,0x95,0x67,0x8E,0xDB,0xC7,0x05,0xE7};

        using (var aes = Aes.Create())
        using (var encryptor = aes.CreateEncryptor(key, key))
        {
            var plainText = Encoding.UTF8.GetBytes(toEncrypt);
            return encryptor.TransformFinalBlock(plainText, 0, plainText.Length);
        }
    }

    public static string DecryptFromBase64ToString(string base64Encrypted)
    {
        byte[] b64Bytes = System.Convert.FromBase64String(base64Encrypted);
        return DecryptToString(b64Bytes);
    }

    public static string DecryptToString(byte[] encryptedData)
    {
        //var key = GetKey(password);
        byte[] key = {0x05,0x12,0x3D,0x2C,0x7D,0x22,0xF7,0x5A,0x9B,0x95,0x67,0x8E,0xDB,0xC7,0x05,0xE7};
        using (var aes = Aes.Create())
        using (var encryptor = aes.CreateDecryptor(key, key))
        {
            var decryptedBytes = encryptor
                .TransformFinalBlock(encryptedData, 0, encryptedData.Length);
            return Encoding.UTF8.GetString(decryptedBytes);
        }
    }

    // converts password to 128 bit hash
    private static byte[] GetKey(string password)
    {
        var keyBytes = Encoding.UTF8.GetBytes(password);
        using (var md5 = MD5.Create())
        {
            return md5.ComputeHash(keyBytes);
        }
    }
}

