import { defineStore } from "pinia";
import { ref } from "vue";

function isObject(value: unknown): value is object {
  const type = typeof value;
  return type === "object" && value !== null;
}

export const SUPPORTED_LOCALES = ["de", "en"] as const;

export type Locale = (typeof SUPPORTED_LOCALES)[number];

const useI18nStore = defineStore("i18n", () => {
  // TODO: can be replaces by useRoute().params.locale <<-- current locale always up to date
  const currentLocale = ref<Locale>();
  const messages: Record<Locale, null | Record<string, unknown>> = { de: null, en: null };

  async function setLocale(newLocale: Locale) {
    document.querySelector("html")?.setAttribute("lang", newLocale);

    if (!messages[newLocale]) {
      let newMessages;
      if (newLocale === "de") {
        newMessages = await import("@/assets/locales/de.json");
      } else if (newLocale === "en") {
        newMessages = await import("@/assets/locales/en.json");
      } else {
        throw new Error(`Unknown requested locale ${newLocale}`);
      }
      messages[newLocale] = newMessages?.default;
    }

    currentLocale.value = newLocale;
    return true;
  }

  function t(messageId: string, data?: Record<string, string | number>) {
    if (currentLocale.value) {
      const localeMessages = messages[currentLocale.value];
      if (!localeMessages) {
        throw new Error(`Locale ${currentLocale.value} not downloaded yet.`);
      }
      // @ts-ignore - reduce does not provide a type annotation that is needed here. reduce<U>((p: U, c:T) => U, initial: U): U is not what I want
      const message = messageId.split(".").reduce((o, key) => {
        if (isObject(o)) return o[key];
      }, messages[currentLocale.value]);

      if (!message) {
        throw new Error(
          `Message ID '${messageId}' is not defined! Set it is on '${currentLocale.value}.json'`
        );
      }

      if (data) {
        return replaceVariables(message, data);
      }

      return message;
    }
    return "";
  }

  return {
    currentLocale,
    setLocale,
    t
  };
});

function replaceVariables(message: string, variables: Record<string, string | number>) {
  return message.replace(/{([^{}]+)}/g, (match, variableName) => {
    return String(variables[variableName]) || match;
  });
}

export { useI18nStore };
