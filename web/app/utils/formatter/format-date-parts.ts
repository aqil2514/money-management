import type { TransactionItemHeader } from "~/types/transaction";
import { id } from "date-fns/locale";
import { format } from "date-fns";
import formatToRupiah from "./format-to-rupiah";

export default (header: TransactionItemHeader) => {
  const dateObj = new Date(header.date);

  const income = formatToRupiah(header.income);
  const expense = formatToRupiah(header.expense);
  return {
    income,
    expense,
    date: format(dateObj, "dd"),
    day: format(dateObj, "EEEE", { locale: id }),
    month: format(dateObj, "MM"),
    year: format(dateObj, "yyyy"),
  };
};
