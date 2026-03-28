import hashlib
import base64

def get_mus_anon_var_name(prefix, data):
    """
    Generates a unique Go identifier for an anonymous MUS serializer.
    
    Args:
        prefix: The type prefix (e.g., 'string', 'map', 'slice').
        data: The deterministic string describing the type and its options.
        
    Returns:
        A unique Go identifier following the project's naming convention.
    """
    # 1. MD5 Hash of the data
    md5_hash = hashlib.md5(data.encode('utf-8')).digest()
    
    # 2. Base64 Encode
    b64_str = base64.b64encode(md5_hash).decode('utf-8')
    
    # 3. Project-specific substitutions
    # + -> Δ (U+0394 Delta)
    # / -> Σ (U+03A3 Sigma)
    # = -> Ξ (U+039E Xi)
    substitutions = {
        '+': 'Δ',
        '/': 'Σ',
        '=': 'Ξ',
    }
    
    for char, replacement in substitutions.items():
        b64_str = b64_str.replace(char, replacement)
        
    return f"{prefix}{b64_str}"

if __name__ == "__main__":
    import sys
    if len(sys.argv) < 3:
        print("Usage: python hash_gen.py <prefix> <data>")
        sys.exit(1)
    
    print(get_mus_anon_var_name(sys.argv[1], sys.argv[2]))
