export default (
  nominal: number,
  options?: {
    withSymbol?: boolean;
    withSign?: boolean;
  },
) => {
  const { withSymbol = true, withSign = false } = options || {};

  const formatter = new Intl.NumberFormat("id-ID", {
    style: withSymbol ? "currency" : "decimal",
    currency: "IDR",
    minimumFractionDigits: 0,
    maximumFractionDigits: 0,
  });

  let result = formatter.format(Math.abs(nominal));

  if (withSign) {
    const sign = nominal > 0 ? "+" : nominal < 0 ? "-" : "";
    return `${sign} ${result}`;
  }

  return nominal < 0 && !withSymbol ? `-${result}` : formatter.format(nominal);
};
