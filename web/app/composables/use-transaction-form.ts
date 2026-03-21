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
    const rawTime: Time = event.data.time;

    const createdDate = new Date(
      rawDate.year,
      rawDate.month,
      rawDate.day,
      rawTime.hour,
      rawTime.minute,
    );

    const payload = {
      ...event.data,
      date: createdDate.toISOString(),
    };

    try {
      const response = await $fetch("http://localhost:8000/transactions", {
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
