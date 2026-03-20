<script setup lang="ts">
import type { CalendarDate } from '@internationalized/date';


const model = defineModel<CalendarDate>()

const inputDate = useTemplateRef('inputDate')
</script>

<template>
    <UFormField label="Tanggal Transaksi" name="date">
        <UInputDate ref="inputDate" v-model="model" class="w-full">
            <template #trailing>
                <UPopover :reference="inputDate?.inputsRef[3]?.$el">
                    <UButton color="neutral" variant="link" size="sm" icon="i-lucide-calendar" class="px-0" />
                    <template #content>
                        <UCalendar v-model="model" class="p-2" :month-controls="true" />
                    </template>
                </UPopover>
            </template>
        </UInputDate>
    </UFormField>
</template>