import { SUPPORTED_LOCALES, type Locale } from '@/stores/i18n';

export function getSupportedBrowserLocale(): Locale | undefined {
  const browserLocale = navigator.language.split('-')[0];
  if (isSupportedLocale(browserLocale)) {
    return browserLocale;
  }
}

export function isSupportedLocale(locale: string): locale is Locale {
  return (SUPPORTED_LOCALES as readonly string[]).includes(locale);
}

export function supportedOrFallbackLocale(locale: string | string[]): Locale {
  if (Array.isArray(locale) || !isSupportedLocale(locale))
    return getSupportedBrowserLocale() || 'en';
  else return locale;
}
