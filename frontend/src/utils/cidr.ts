// function used to check if a given string is a valid CIDR
export default function isCIDR(cidr: string): boolean {
  const cidrRegex = new RegExp(
    /^(\d{1,3}\.){3}\d{1,3}\/\d{1,2}$/
  );
  return cidrRegex.test(cidr);
}
