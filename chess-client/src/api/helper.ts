import { handlePromise } from "kas-helpers";

export const jsonFetch = async <T>(url: string, options: RequestInit = {}) => {
    const { result } = await handlePromise(fetch(url, {
        headers: {
        "Content-Type": "application/json",
        },
        ...options,
    }), (result) => !result.ok);

    if (result) return await result.json() as T;

    throw new Error("fail to call endpoint");
}