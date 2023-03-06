import { Options } from "../types/options";

export const getConfig = (): Options => {
    return {
        defaultRegion: "us-west-2",
        apiStageName: "main",
    }
}
