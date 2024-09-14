declare global {
    namespace NodeJS {
        interface ProcessEnv {
            [key: string]: string | undefined;
            PRIVATE_TEST_KEY: string;
            PRIVATE_KEY: string;
        }
    }
}

export {}