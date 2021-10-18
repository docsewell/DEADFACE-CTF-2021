#include <stdio.h>
#include <string.h>

int main() {

    int index[33] = {18, 23, 1, 0, 19, 17, 15, 31, 11, 30, 10, 5, 21, 13, 20, 16, 14, 25, 8, 4, 3, 27, 7, 9, 2, 22, 6, 28, 24, 29, 32, 12, 26};
    int index2[18] = {8, 2, 3, 7, 5, 14, 10, 1, 4, 6, 0, 11, 9, 17, 12, 13, 16, 15};

    //char* passphrase = "B00-Boo-Boo-B33ry!"; /* 18 characters */
    //char* true_flag = "flag{B00-B00-B00-Bury-IZ-DA-BOMB}"; /* 33 characters *//
    // 3 good characters, followed by 3 garbage characters
    char* obs_passphrase =
            "B0-"
            "z%\xDD"
            "-o3"
            "3%\xEE"
            "o0B"
            "9)\xEE"
            "oB-"
            "@3\xDD"
            "o!B"
            "8(\xEE"
            "3yr"
            "!!\xEE"; /* 18 characters */

    // 11 good characters, followed by 11 garbage characters
    char* obs_flag =
            "uZlfrB0B0M0"
            "THE\xDDZEALOT-"
            "B-By-0D-{g-"
            "FL\x06G{Captai"
            "0BaI0B-O}-A"
            "n-Crunch}Ze";

    char password[512]; // input buffer
    char true_flag[34];
    char almost_true_flag[34];
    char almost_true_passphrase[19];
    char true_passphrase[19];

    char* halloweenie_message = "What is the best and sp00kiest breakfast cereal?";
    char* password_prompt = "Please enter the passphrase: ";
    char* not_flag = "notflag{you-guessed-it-again--this-is-not-the-flag}";

    // Deobfuscate the passphrase
    {
        int move = 0;
        int x = 0;
        int y = 0;
        while (x < 36) {
            if (x % 3 == 0) {
                move = !move;
            }
            if (move) {
                almost_true_passphrase[y] = obs_passphrase[x];
                y++;
            }
            x++;
        }

        almost_true_passphrase[18] = '\x00';
        int j = 0;
        for (int i = 0; i < 18; i++) {
            // Find the index of i in the array called "index"
            for (j = 0; i < 18; j++) {
                if (index2[j] == i) {
                    break;
                }
            }
            true_passphrase[i] = almost_true_passphrase[j];
        }
        true_passphrase[18] = '\x00';
    }

    // Deobfuscate the flag
    // Block for re-using variable names
    {
        int move = 0;
        int x = 0;
        int y = 0;
        while (x < 66) {
            if (x % 11 == 0) {
                move = !move;
            }
            if (move) {
                almost_true_flag[y] = obs_flag[x];
                y++;
            }
            x++;
        }

        almost_true_flag[33] = '\x00';
        int j = 0;
        for (int i = 0; i < 33; i++) {
            // Find the index of i in the array called "index"
            for (j = 0; i < 33; j++) {
                if (index[j] == i) {
                    break;
                }
            }
            true_flag[i] = almost_true_flag[j];
        }
        true_flag[33] = '\x00';

    }

    printf("%s\n", halloweenie_message);
    printf("%s", password_prompt);
    scanf("%s", password);

    int res = 0;        // result of functions
    res = strcmp(password, true_passphrase);

    if (res == 0) {
        printf("%s\n", true_flag);
    } else {
        printf("%s\n", not_flag);
    }

    return 0;
}


