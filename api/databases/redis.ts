import { Redis } from "ioredis";

export const redisClient = new Redis("localhost:6379");
