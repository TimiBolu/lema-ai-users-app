import { z } from "zod";

export const EnvConfigSchema = z.object({
  apiBaseUrl: z.string(),
  apiEnv: z.string(),
});

type EnvConfig = z.infer<typeof EnvConfigSchema>;
const EnvConfig = {
  apiBaseUrl: import.meta.env.VITE_APP_API_BASE_URL,
  apiEnv: import.meta.env.VITE_APP_ENV,
} as EnvConfig;

EnvConfigSchema.parse(EnvConfig);

export default EnvConfig;
