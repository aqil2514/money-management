// composables/useTransactionForm.ts
import * as z from "zod";
import { CalendarDate, getLocalTimeZone, Time } from "@internationalized/date";
import type { FormSubmitEvent } from "@nuxt/ui";
import {
  defaultTransaction,
  type TransactionSchemaType,
} from "~/schemas/transaction-schema";

export const useTransactionForm = () => {
  const toast = useToast();

  const now = new Date();

  const state = reactive({
    ...defaultTransaction,
    time: new Time(now.getHours(), now.getMinutes()),
  }) as TransactionSchemaType;

  async function onSubmit(event: FormSubmitEvent<TransactionSchemaType>) {
    const rawDate: CalendarDate = event.data.date;
    const formattedDate = rawDate.toDate(getLocalTimeZone());

    const rawTime: Time = event.data.time;
    const formattedTime = rawTime.toString();

    const payload = {
      ...event.data,
      date: formattedDate,
      time: formattedTime,
    };

    console.log(payload);

    return;
    try {
      const response = await $fetch("http://localhost:8000/api/transactions", {
        method: "POST",
        body: payload,
      });

      toast.add({
        title: "Berhasil!",
        description: "Data transaksi telah tersimpan di server Go.",
        color: "success",
      });

      console.log("Respon dari Go:", response);
    } catch (error: any) {
      console.error(error);
      toast.add({
        title: "Gagal!",
        description:
          error.data?.message || "Tidak dapat terhubung ke server Go.",
        color: "error",
      });
    }
  }

  return {
    state,
    onSubmit,
  };
};
