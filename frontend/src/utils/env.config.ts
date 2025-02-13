import { z } from "zod";

export const EnvConfigSchema = z.object({
  apiBaseUrl: z.string(),
  apiEnv: z.string(),
  jwtToken: z.string(),
});

type EnvConfig = z.infer<typeof EnvConfigSchema>;
const EnvConfig = {
  apiBaseUrl: import.meta.env.VITE_APP_API_BASE_URL,
  apiEnv: import.meta.env.VITE_APP_ENV,
  jwtToken: import.meta.env.VITE_JWT_TOKEN,
} as EnvConfig;

EnvConfigSchema.parse(EnvConfig);

export default EnvConfig;
