import { redisClient } from "../databases/redis";

export async function getFixturesFromCache(page: number, limit: number): Promise<any> {
    return new Promise((resolve, reject) => {
      redisClient.keys("*", async (err, keys) => {
        if (err) {
          return reject(err);
        }
  
        if (!keys || keys.length === 0) {
          return resolve(null);
        }
  
        const cachedDataPromises: Promise<string | null>[] = keys.map((key: string) => {
          return new Promise((resolve, reject) => {
            redisClient.get(key, (err, cachedData) => {
              if (err) {
                reject(err);
              } else {
                resolve(cachedData || null);
              }
            });
          });
        });
        
        Promise.all(cachedDataPromises)
          .then((cachedDataArray: (string | null)[]) => {
            const combinedCachedData: any[] = cachedDataArray.reduce((acc: any[], cachedData: string | null) => {
              if (cachedData) {
                const parsedData = JSON.parse(cachedData);
                return acc.concat(parsedData);
              } else {
                return acc;
              }
            }, []);
  
            const startIndex: number = (page - 1) * limit;
            const endIndex: number = page * limit;
            const fixtures = combinedCachedData.slice(startIndex, endIndex);
            const totalPages: number = Math.ceil(combinedCachedData.length / limit);
  
            resolve({ fixtures, totalPages });
          })
          .catch(reject);
      });
    });
  }