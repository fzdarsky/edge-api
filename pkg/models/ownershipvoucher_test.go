package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"encoding/json"

	"github.com/redhatinsights/edge-api/pkg/models"
	ovde "github.com/redhatinsights/edge-api/pkg/ownershipvoucher"
)

var (
	edgeDeviceName string = "testdevice1"
	edgeDeviceUUID string = "4609ec7b-b56d-4f9f-c62c-0e73f513ec55"
	ovb = []byte{132, 153, 3, 232, 24, 134, 24, 24, 24, 100, 24, 144, 24, 24, 24, 70, 9, 24, 24, 24, 236, 24, 24, 24, 123, 24, 24, 24, 181, 24, 24, 24, 109, 24, 24, 24, 79, 24, 24, 24, 159, 24, 24, 24, 198, 24, 24, 24, 44, 14, 24, 24, 24, 115, 24, 24, 24, 245, 19, 24, 24, 24, 236, 24, 24, 24, 85, 24, 129, 24, 132, 24, 130, 5, 24, 105, 24, 49, 24, 48, 24, 46, 24, 56, 24, 57, 24, 46, 24, 48, 24, 46, 24, 51, 24, 130, 3, 24, 25, 24, 31, 24, 145, 24, 130, 4, 24, 25, 24, 31, 24, 145, 24, 130, 12, 1, 24, 107, 24, 116, 24, 101, 24, 115, 24, 116, 24, 100, 24, 101, 24, 118, 24, 105, 24, 99, 24, 101, 24, 49, 24, 131, 24, 38, 1, 24, 153, 1, 24, 219, 24, 24, 24, 48, 24, 24, 24, 130, 1, 24, 24, 24, 215, 24, 24, 24, 48, 24, 24, 24, 130, 1, 24, 24, 24, 125, 24, 24, 24, 160, 3, 2, 1, 2, 2, 20, 24, 24, 24, 121, 24, 24, 24, 117, 24, 24, 24, 174, 24, 24, 24, 62, 24, 24, 24, 135, 24, 24, 24, 249, 24, 24, 24, 236, 24, 24, 24, 197, 24, 24, 24, 187, 24, 24, 24, 191, 24, 24, 24, 219, 24, 24, 24, 169, 24, 24, 24, 36, 24, 24, 24, 123, 24, 24, 24, 187, 24, 24, 24, 226, 10, 24, 24, 24, 125, 24, 24, 24, 116, 24, 24, 24, 25, 24, 24, 24, 48, 10, 6, 8, 24, 24, 24, 42, 24, 24, 24, 134, 24, 24, 24, 72, 24, 24, 24, 206, 24, 24, 24, 61, 4, 3, 2, 24, 24, 24, 48, 24, 24, 24, 65, 24, 24, 24, 49, 11, 24, 24, 24, 48, 9, 6, 3, 24, 24, 24, 85, 4, 6, 19, 2, 24, 24, 24, 85, 24, 24, 24, 83, 24, 24, 24, 49, 22, 24, 24, 24, 48, 20, 6, 3, 24, 24, 24, 85, 4, 10, 12, 13, 24, 24, 24, 82, 24, 24, 24, 72, 24, 24, 24, 69, 24, 24, 24, 76, 24, 24, 24, 32, 24, 24, 24, 102, 24, 24, 24, 111, 24, 24, 24, 114, 24, 24, 24, 32, 24, 24, 24, 69, 24, 24, 24, 100, 24, 24, 24, 103, 24, 24, 24, 101, 24, 24, 24, 49, 24, 24, 24, 26, 24, 24, 24, 48, 24, 24, 24, 24, 6, 3, 24, 24, 24, 85, 4, 3, 12, 17, 24, 24, 24, 70, 24, 24, 24, 73, 24, 24, 24, 68, 24, 24, 24, 79, 24, 24, 24, 32, 24, 24, 24, 77, 24, 24, 24, 97, 24, 24, 24, 110, 24, 24, 24, 117, 24, 24, 24, 102, 24, 24, 24, 97, 24, 24, 24, 99, 24, 24, 24, 116, 24, 24, 24, 117, 24, 24, 24, 114, 24, 24, 24, 101, 24, 24, 24, 114, 24, 24, 24, 48, 24, 24, 24, 30, 23, 13, 24, 24, 24, 50, 24, 24, 24, 49, 24, 24, 24, 49, 24, 24, 24, 48, 24, 24, 24, 49, 24, 24, 24, 52, 24, 24, 24, 49, 24, 24, 24, 49, 24, 24, 24, 51, 24, 24, 24, 48, 24, 24, 24, 48, 24, 24, 24, 52, 24, 24, 24, 90, 23, 13, 24, 24, 24, 50, 24, 24, 24, 50, 24, 24, 24, 49, 24, 24, 24, 48, 24, 24, 24, 49, 24, 24, 24, 52, 24, 24, 24, 49, 24, 24, 24, 49, 24, 24, 24, 51, 24, 24, 24, 48, 24, 24, 24, 48, 24, 24, 24, 52, 24, 24, 24, 90, 24, 24, 24, 48, 24, 24, 24, 65, 24, 24, 24, 49, 11, 24, 24, 24, 48, 9, 6, 3, 24, 24, 24, 85, 4, 6, 19, 2, 24, 24, 24, 85, 24, 24, 24, 83, 24, 24, 24, 49, 22, 24, 24, 24, 48, 20, 6, 3, 24, 24, 24, 85, 4, 10, 12, 13, 24, 24, 24, 82, 24, 24, 24, 72, 24, 24, 24, 69, 24, 24, 24, 76, 24, 24, 24, 32, 24, 24, 24, 102, 24, 24, 24, 111, 24, 24, 24, 114, 24, 24, 24, 32, 24, 24, 24, 69, 24, 24, 24, 100, 24, 24, 24, 103, 24, 24, 24, 101, 24, 24, 24, 49, 24, 24, 24, 26, 24, 24, 24, 48, 24, 24, 24, 24, 6, 3, 24, 24, 24, 85, 4, 3, 12, 17, 24, 24, 24, 70, 24, 24, 24, 73, 24, 24, 24, 68, 24, 24, 24, 79, 24, 24, 24, 32, 24, 24, 24, 77, 24, 24, 24, 97, 24, 24, 24, 110, 24, 24, 24, 117, 24, 24, 24, 102, 24, 24, 24, 97, 24, 24, 24, 99, 24, 24, 24, 116, 24, 24, 24, 117, 24, 24, 24, 114, 24, 24, 24, 101, 24, 24, 24, 114, 24, 24, 24, 48, 24, 24, 24, 89, 24, 24, 24, 48, 19, 6, 7, 24, 24, 24, 42, 24, 24, 24, 134, 24, 24, 24, 72, 24, 24, 24, 206, 24, 24, 24, 61, 2, 1, 6, 8, 24, 24, 24, 42, 24, 24, 24, 134, 24, 24, 24, 72, 24, 24, 24, 206, 24, 24, 24, 61, 3, 1, 7, 3, 24, 24, 24, 66, 0, 4, 24, 24, 24, 183, 24, 24, 24, 25, 24, 24, 24, 131, 24, 24, 24, 140, 24, 24, 24, 90, 24, 24, 24, 103, 24, 24, 24, 126, 24, 24, 24, 226, 24, 24, 24, 164, 24, 24, 24, 233, 24, 24, 24, 84, 24, 24, 24, 226, 24, 24, 24, 34, 24, 24, 24, 211, 24, 24, 24, 143, 24, 24, 24, 93, 24, 24, 24, 101, 24, 24, 24, 77, 24, 24, 24, 223, 24, 24, 24, 141, 24, 24, 24, 47, 24, 24, 24, 42, 24, 24, 24, 198, 24, 24, 24, 250, 24, 24, 24, 77, 24, 24, 24, 52, 24, 24, 24, 90, 24, 24, 24, 128, 24, 24, 24, 131, 13, 24, 24, 24, 28, 24, 24, 24, 239, 24, 24, 24, 240, 8, 24, 24, 24, 163, 24, 24, 24, 148, 8, 24, 24, 24, 77, 24, 24, 24, 151, 24, 24, 24, 195, 24, 24, 24, 120, 24, 24, 24, 64, 24, 24, 24, 119, 24, 24, 24, 84, 24, 24, 24, 245, 24, 24, 24, 68, 24, 24, 24, 167, 24, 24, 24, 64, 24, 24, 24, 164, 24, 24, 24, 250, 24, 24, 24, 225, 24, 24, 24, 233, 24, 24, 24, 84, 7, 24, 24, 24, 30, 22, 24, 24, 24, 227, 24, 24, 24, 197, 24, 24, 24, 121, 24, 24, 24, 106, 24, 24, 24, 214, 24, 24, 24, 74, 24, 24, 24, 147, 24, 24, 24, 41, 24, 24, 24, 163, 24, 24, 24, 83, 24, 24, 24, 48, 24, 24, 24, 81, 24, 24, 24, 48, 24, 24, 24, 29, 6, 3, 24, 24, 24, 85, 24, 24, 24, 29, 14, 4, 22, 4, 20, 24, 24, 24, 188, 24, 24, 24, 109, 7, 24, 24, 24, 35, 24, 24, 24, 30, 24, 24, 24, 98, 24, 24, 24, 133, 3, 24, 24, 24, 246, 24, 24, 24, 115, 21, 24, 24, 24, 103, 24, 24, 24, 133, 24, 24, 24, 73, 24, 24, 24, 160, 24, 24, 24, 36, 24, 24, 24, 249, 24, 24, 24, 133, 24, 24, 24, 206, 24, 24, 24, 117, 24, 24, 24, 48, 24, 24, 24, 31, 6, 3, 24, 24, 24, 85, 24, 24, 24, 29, 24, 24, 24, 35, 4, 24, 24, 24, 24, 24, 24, 24, 48, 22, 24, 24, 24, 128, 20, 24, 24, 24, 188, 24, 24, 24, 109, 7, 24, 24, 24, 35, 24, 24, 24, 30, 24, 24, 24, 98, 24, 24, 24, 133, 3, 24, 24, 24, 246, 24, 24, 24, 115, 21, 24, 24, 24, 103, 24, 24, 24, 133, 24, 24, 24, 73, 24, 24, 24, 160, 24, 24, 24, 36, 24, 24, 24, 249, 24, 24, 24, 133, 24, 24, 24, 206, 24, 24, 24, 117, 24, 24, 24, 48, 15, 6, 3, 24, 24, 24, 85, 24, 24, 24, 29, 19, 1, 1, 24, 24, 24, 255, 4, 5, 24, 24, 24, 48, 3, 1, 1, 24, 24, 24, 255, 24, 24, 24, 48, 10, 6, 8, 24, 24, 24, 42, 24, 24, 24, 134, 24, 24, 24, 72, 24, 24, 24, 206, 24, 24, 24, 61, 4, 3, 2, 3, 24, 24, 24, 72, 0, 24, 24, 24, 48, 24, 24, 24, 69, 2, 24, 24, 24, 32, 24, 24, 24, 59, 24, 24, 24, 134, 24, 24, 24, 30, 24, 24, 24, 182, 24, 24, 24, 233, 24, 24, 24, 237, 24, 24, 24, 141, 24, 24, 24, 250, 24, 24, 24, 90, 24, 24, 24, 97, 24, 24, 24, 230, 24, 24, 24, 199, 24, 24, 24, 30, 24, 24, 24, 142, 24, 24, 24, 137, 24, 24, 24, 188, 24, 24, 24, 54, 24, 24, 24, 199, 24, 24, 24, 132, 24, 24, 24, 173, 24, 24, 24, 240, 24, 24, 24, 200, 24, 24, 24, 75, 24, 24, 24, 151, 24, 24, 24, 129, 24, 24, 24, 157, 24, 24, 24, 149, 24, 24, 24, 63, 24, 24, 24, 40, 24, 24, 24, 60, 24, 24, 24, 253, 24, 24, 24, 204, 2, 24, 24, 24, 33, 0, 24, 24, 24, 177, 24, 24, 24, 175, 24, 24, 24, 206, 24, 24, 24, 192, 24, 24, 24, 53, 24, 24, 24, 30, 24, 24, 24, 85, 24, 24, 24, 207, 24, 24, 24, 40, 15, 24, 24, 24, 217, 24, 24, 24, 173, 11, 24, 24, 24, 147, 24, 24, 24, 48, 24, 24, 24, 152, 24, 24, 24, 182, 24, 24, 24, 190, 3, 24, 24, 24, 199, 24, 24, 24, 214, 24, 24, 24, 109, 24, 24, 24, 167, 24, 24, 24, 50, 24, 24, 24, 188, 24, 24, 24, 238, 24, 24, 24, 197, 24, 24, 24, 222, 24, 24, 24, 78, 24, 24, 24, 167, 8, 24, 24, 24, 173, 24, 130, 24, 56, 24, 42, 24, 152, 24, 48, 24, 24, 24, 113, 24, 24, 24, 109, 7, 24, 24, 24, 186, 24, 24, 24, 221, 24, 24, 24, 65, 24, 24, 24, 134, 24, 24, 24, 57, 24, 24, 24, 169, 24, 24, 24, 112, 24, 24, 24, 148, 24, 24, 24, 212, 24, 24, 24, 229, 24, 24, 24, 102, 24, 24, 24, 145, 24, 24, 24, 43, 24, 24, 24, 48, 24, 24, 24, 162, 24, 24, 24, 243, 24, 24, 24, 221, 24, 24, 24, 205, 24, 24, 24, 250, 24, 24, 24, 59, 24, 24, 24, 146, 24, 24, 24, 245, 24, 24, 24, 236, 24, 24, 24, 52, 24, 24, 24, 171, 24, 24, 24, 106, 0, 24, 24, 24, 91, 24, 24, 24, 146, 10, 18, 24, 24, 24, 148, 24, 24, 24, 164, 24, 24, 24, 170, 24, 24, 24, 192, 24, 24, 24, 72, 24, 24, 24, 154, 24, 24, 24, 35, 24, 24, 24, 106, 24, 24, 24, 150, 24, 24, 24, 144, 24, 24, 24, 140, 24, 24, 24, 140, 24, 24, 24, 67, 24, 24, 24, 83, 130, 56, 42, 152, 48, 23, 24, 124, 17, 24, 165, 24, 49, 24, 51, 24, 134, 24, 147, 24, 122, 24, 107, 24, 70, 24, 128, 24, 156, 24, 232, 24, 247, 24, 200, 24, 40, 24, 158, 24, 218, 24, 206, 24, 151, 24, 116, 24, 197, 24, 203, 24, 170, 24, 199, 24, 121, 24, 49, 24, 168, 24, 66, 24, 127, 24, 146, 24, 160, 24, 227, 24, 175, 24, 119, 24, 152, 24, 111, 24, 110, 17, 21, 16, 24, 177, 24, 100, 24, 162, 24, 203, 24, 26, 24, 195, 130, 153, 1, 68, 24, 48, 24, 130, 1, 24, 64, 24, 48, 24, 129, 24, 231, 24, 160, 3, 2, 1, 2, 2, 9, 0, 24, 203, 6, 24, 160, 24, 119, 24, 172, 24, 70, 24, 72, 24, 33, 24, 48, 10, 6, 8, 24, 42, 24, 134, 24, 72, 24, 206, 24, 61, 4, 3, 3, 24, 48, 24, 54, 24, 49, 11, 24, 48, 9, 6, 3, 24, 85, 4, 6, 19, 2, 24, 85, 24, 83, 24, 49, 22, 24, 48, 20, 6, 3, 24, 85, 4, 10, 12, 13, 24, 82, 24, 72, 24, 69, 24, 76, 24, 32, 24, 102, 24, 111, 24, 114, 24, 32, 24, 69, 24, 100, 24, 103, 24, 101, 24, 49, 15, 24, 48, 13, 6, 3, 24, 85, 4, 3, 12, 6, 24, 68, 24, 101, 24, 118, 24, 105, 24, 99, 24, 101, 24, 48, 24, 30, 23, 13, 24, 50, 24, 49, 24, 49, 24, 48, 24, 49, 24, 52, 24, 49, 24, 49, 24, 52, 24, 48, 24, 53, 24, 52, 24, 90, 23, 13, 24, 51, 24, 49, 24, 49, 24, 48, 24, 49, 24, 50, 24, 49, 24, 49, 24, 52, 24, 48, 24, 53, 24, 52, 24, 90, 24, 48, 22, 24, 49, 20, 24, 48, 18, 6, 3, 24, 85, 4, 3, 12, 11, 24, 116, 24, 101, 24, 115, 24, 116, 24, 100, 24, 101, 24, 118, 24, 105, 24, 99, 24, 101, 24, 49, 24, 48, 24, 89, 24, 48, 19, 6, 7, 24, 42, 24, 134, 24, 72, 24, 206, 24, 61, 2, 1, 6, 8, 24, 42, 24, 134, 24, 72, 24, 206, 24, 61, 3, 1, 7, 3, 24, 66, 0, 4, 24, 165, 24, 27, 24, 101, 24, 170, 18, 24, 157, 24, 224, 24, 185, 24, 204, 24, 147, 24, 174, 9, 24, 214, 24, 223, 24, 109, 24, 226, 24, 39, 24, 191, 24, 186, 24, 240, 24, 106, 24, 131, 24, 250, 24, 188, 24, 117, 24, 69, 24, 183, 24, 207, 24, 116, 24, 160, 24, 88, 24, 222, 24, 136, 19, 24, 219, 24, 210, 24, 123, 2, 24, 105, 24, 136, 10, 24, 72, 24, 72, 24, 223, 24, 38, 24, 117, 24, 55, 24, 205, 24, 147, 24, 88, 24, 50, 24, 167, 24, 45, 24, 38, 24, 139, 6, 24, 150, 24, 250, 24, 161, 24, 95, 24, 45, 24, 228, 24, 230, 24, 135, 24, 48, 10, 6, 8, 24, 42, 24, 134, 24, 72, 24, 206, 24, 61, 4, 3, 3, 3, 24, 72, 0, 24, 48, 24, 69, 2, 24, 32, 24, 119, 24, 123, 24, 169, 24, 248, 24, 94, 24, 174, 24, 89, 24, 169, 24, 48, 24, 194, 24, 104, 24, 53, 24, 242, 24, 139, 24, 207, 24, 206, 24, 86, 24, 174, 24, 151, 24, 117, 24, 239, 24, 76, 24, 204, 24, 84, 24, 173, 24, 141, 24, 228, 24, 56, 24, 246, 24, 185, 24, 100, 24, 253, 2, 24, 33, 0, 24, 130, 24, 177, 24, 155, 24, 93, 21, 24, 25, 24, 126, 24, 95, 24, 143, 24, 37, 24, 43, 24, 88, 24, 187, 24, 177, 24, 249, 24, 234, 18, 24, 131, 24, 105, 24, 101, 24, 88, 24, 127, 24, 169, 24, 164, 24, 175, 24, 93, 24, 145, 24, 220, 24, 39, 24, 201, 24, 186, 24, 131, 153, 1, 197, 24, 48, 24, 130, 1, 24, 193, 24, 48, 24, 130, 1, 24, 103, 24, 160, 3, 2, 1, 2, 2, 20, 24, 119, 24, 168, 24, 149, 24, 216, 24, 113, 24, 100, 24, 109, 24, 78, 24, 81, 24, 138, 24, 205, 24, 243, 24, 112, 24, 170, 24, 26, 24, 186, 24, 156, 24, 43, 24, 140, 24, 77, 24, 48, 10, 6, 8, 24, 42, 24, 134, 24, 72, 24, 206, 24, 61, 4, 3, 2, 24, 48, 24, 54, 24, 49, 11, 24, 48, 9, 6, 3, 24, 85, 4, 6, 19, 2, 24, 85, 24, 83, 24, 49, 22, 24, 48, 20, 6, 3, 24, 85, 4, 10, 12, 13, 24, 82, 24, 72, 24, 69, 24, 76, 24, 32, 24, 102, 24, 111, 24, 114, 24, 32, 24, 69, 24, 100, 24, 103, 24, 101, 24, 49, 15, 24, 48, 13, 6, 3, 24, 85, 4, 3, 12, 6, 24, 68, 24, 101, 24, 118, 24, 105, 24, 99, 24, 101, 24, 48, 24, 30, 23, 13, 24, 50, 24, 49, 24, 49, 24, 48, 24, 49, 24, 52, 24, 49, 24, 49, 24, 51, 24, 48, 24, 48, 24, 52, 24, 90, 23, 13, 24, 50, 24, 50, 24, 49, 24, 48, 24, 49, 24, 52, 24, 49, 24, 49, 24, 51, 24, 48, 24, 48, 24, 52, 24, 90, 24, 48, 24, 54, 24, 49, 11, 24, 48, 9, 6, 3, 24, 85, 4, 6, 19, 2, 24, 85, 24, 83, 24, 49, 22, 24, 48, 20, 6, 3, 24, 85, 4, 10, 12, 13, 24, 82, 24, 72, 24, 69, 24, 76, 24, 32, 24, 102, 24, 111, 24, 114, 24, 32, 24, 69, 24, 100, 24, 103, 24, 101, 24, 49, 15, 24, 48, 13, 6, 3, 24, 85, 4, 3, 12, 6, 24, 68, 24, 101, 24, 118, 24, 105, 24, 99, 24, 101, 24, 48, 24, 89, 24, 48, 19, 6, 7, 24, 42, 24, 134, 24, 72, 24, 206, 24, 61, 2, 1, 6, 8, 24, 42, 24, 134, 24, 72, 24, 206, 24, 61, 3, 1, 7, 3, 24, 66, 0, 4, 24, 29, 24, 212, 24, 153, 24, 130, 24, 233, 17, 24, 234, 24, 106, 24, 82, 24, 255, 24, 100, 24, 117, 24, 58, 24, 165, 24, 228, 7, 24, 228, 24, 196, 24, 193, 24, 97, 24, 110, 18, 24, 202, 24, 130, 24, 140, 24, 237, 24, 113, 24, 244, 24, 115, 24, 95, 24, 179, 24, 190, 24, 190, 24, 219, 24, 188, 24, 80, 24, 223, 24, 115, 24, 86, 20, 24, 209, 24, 92, 24, 247, 24, 172, 24, 223, 24, 85, 24, 198, 0, 24, 152, 24, 105, 24, 240, 24, 31, 24, 232, 24, 219, 24, 30, 2, 24, 230, 24, 234, 24, 74, 24, 230, 24, 62, 24, 148, 24, 158, 24, 41, 24, 163, 24, 83, 24, 48, 24, 81, 24, 48, 24, 29, 6, 3, 24, 85, 24, 29, 14, 4, 22, 4, 20, 15, 7, 24, 192, 24, 204, 24, 253, 24, 76, 24, 79, 24, 108, 4, 24, 243, 24, 97, 24, 219, 24, 193, 24, 242, 24, 65, 18, 24, 244, 24, 218, 24, 41, 16, 24, 48, 24, 31, 6, 3, 24, 85, 24, 29, 24, 35, 4, 24, 24, 24, 48, 22, 24, 128, 20, 15, 7, 24, 192, 24, 204, 24, 253, 24, 76, 24, 79, 24, 108, 4, 24, 243, 24, 97, 24, 219, 24, 193, 24, 242, 24, 65, 18, 24, 244, 24, 218, 24, 41, 16, 24, 48, 15, 6, 3, 24, 85, 24, 29, 19, 1, 1, 24, 255, 4, 5, 24, 48, 3, 1, 1, 24, 255, 24, 48, 10, 6, 8, 24, 42, 24, 134, 24, 72, 24, 206, 24, 61, 4, 3, 2, 3, 24, 72, 0, 24, 48, 24, 69, 2, 24, 33, 0, 24, 194, 22, 24, 215, 24, 162, 24, 83, 24, 128, 1, 24, 219, 24, 221, 24, 118, 24, 78, 24, 92, 24, 249, 24, 99, 24, 60, 24, 248, 24, 52, 24, 114, 24, 255, 24, 233, 2, 24, 92, 24, 178, 24, 104, 24, 134, 14, 23, 14, 24, 218, 24, 82, 24, 70, 24, 198, 2, 24, 32, 24, 81, 24, 181, 24, 192, 24, 101, 16, 24, 230, 24, 204, 24, 252, 24, 39, 24, 119, 24, 114, 0, 24, 223, 24, 158, 24, 208, 24, 163, 24, 112, 24, 194, 19, 24, 248, 24, 74, 24, 130, 21, 24, 215, 24, 231, 24, 60, 24, 84, 24, 189, 24, 176, 24, 190, 24, 171, 24, 236, 130, 153, 4, 32, 24, 210, 24, 132, 24, 67, 24, 161, 1, 24, 38, 24, 160, 24, 89, 3, 24, 212, 24, 131, 24, 130, 24, 56, 24, 42, 24, 152, 24, 48, 24, 24, 24, 203, 24, 24, 24, 130, 24, 24, 24, 112, 24, 24, 24, 63, 24, 24, 24, 31, 24, 24, 24, 176, 24, 24, 24, 188, 24, 24, 24, 150, 18, 12, 24, 24, 24, 71, 24, 24, 24, 142, 24, 24, 24, 73, 24, 24, 24, 231, 24, 24, 24, 154, 24, 24, 24, 97, 24, 24, 24, 225, 24, 24, 24, 174, 24, 24, 24, 28, 24, 24, 24, 227, 24, 24, 24, 130, 24, 24, 24, 64, 24, 24, 24, 116, 24, 24, 24, 108, 24, 24, 24, 168, 24, 24, 24, 132, 24, 24, 24, 213, 24, 24, 24, 137, 24, 24, 24, 24, 24, 24, 24, 152, 24, 24, 24, 223, 24, 24, 24, 189, 24, 24, 24, 183, 12, 24, 24, 24, 111, 24, 24, 24, 28, 24, 24, 24, 104, 24, 24, 24, 211, 24, 24, 24, 118, 24, 24, 24, 34, 24, 24, 24, 232, 24, 24, 24, 116, 24, 24, 24, 64, 0, 24, 24, 24, 163, 24, 24, 24, 75, 24, 24, 24, 110, 24, 24, 24, 106, 24, 130, 24, 56, 24, 42, 24, 152, 24, 48, 24, 24, 24, 100, 24, 24, 24, 109, 24, 24, 24, 185, 24, 24, 24, 81, 24, 24, 24, 179, 24, 24, 24, 99, 24, 24, 24, 166, 24, 24, 24, 229, 12, 24, 24, 24, 227, 24, 24, 24, 245, 24, 24, 24, 145, 6, 24, 24, 24, 54, 24, 24, 24, 103, 24, 24, 24, 44, 24, 24, 24, 29, 24, 24, 24, 185, 24, 24, 24, 168, 24, 24, 24, 33, 24, 24, 24, 120, 24, 24, 24, 236, 24, 24, 24, 117, 24, 24, 24, 252, 19, 24, 24, 24, 253, 24, 24, 24, 154, 24, 24, 24, 235, 24, 24, 24, 70, 24, 24, 24, 89, 24, 24, 24, 151, 24, 24, 24, 159, 24, 24, 24, 241, 24, 24, 24, 247, 24, 24, 24, 43, 24, 24, 24, 243, 24, 24, 24, 230, 24, 24, 24, 130, 24, 24, 24, 80, 24, 24, 24, 126, 24, 24, 24, 101, 24, 24, 24, 86, 24, 24, 24, 35, 21, 24, 24, 24, 31, 24, 24, 24, 175, 24, 24, 24, 110, 24, 24, 24, 137, 24, 131, 24, 38, 1, 24, 153, 1, 24, 202, 24, 24, 24, 48, 24, 24, 24, 130, 1, 24, 24, 24, 198, 24, 24, 24, 48, 24, 24, 24, 130, 1, 24, 24, 24, 107, 24, 24, 24, 160, 3, 2, 1, 2, 2, 20, 0, 24, 24, 24, 183, 24, 24, 24, 172, 24, 24, 24, 113, 24, 24, 24, 228, 24, 24, 24, 131, 24, 24, 24, 168, 24, 24, 24, 124, 24, 24, 24, 158, 24, 24, 24, 245, 24, 24, 24, 32, 24, 24, 24, 171, 24, 24, 24, 43, 24, 24, 24, 180, 24, 24, 24, 114, 24, 24, 24, 68, 24, 24, 24, 137, 24, 24, 24, 145, 24, 24, 24, 128, 24, 24, 24, 254, 24, 24, 24, 48, 10, 6, 8, 24, 24, 24, 42, 24, 24, 24, 134, 24, 24, 24, 72, 24, 24, 24, 206, 24, 24, 24, 61, 4, 3, 2, 24, 24, 24, 48, 24, 24, 24, 56, 24, 24, 24, 49, 11, 24, 24, 24, 48, 9, 6, 3, 24, 24, 24, 85, 4, 6, 19, 2, 24, 24, 24, 85, 24, 24, 24, 83, 24, 24, 24, 49, 22, 24, 24, 24, 48, 20, 6, 3, 24, 24, 24, 85, 4, 10, 12, 13, 24, 24, 24, 82, 24, 24, 24, 72, 24, 24, 24, 69, 24, 24, 24, 76, 24, 24, 24, 32, 24, 24, 24, 102, 24, 24, 24, 111, 24, 24, 24, 114, 24, 24, 24, 32, 24, 24, 24, 69, 24, 24, 24, 100, 24, 24, 24, 103, 24, 24, 24, 101, 24, 24, 24, 49, 17, 24, 24, 24, 48, 15, 6, 3, 24, 24, 24, 85, 4, 3, 12, 8, 24, 24, 24, 82, 24, 24, 24, 101, 24, 24, 24, 115, 24, 24, 24, 101, 24, 24, 24, 108, 24, 24, 24, 108, 24, 24, 24, 101, 24, 24, 24, 114, 24, 24, 24, 48, 24, 24, 24, 30, 23, 13, 24, 24, 24, 50, 24, 24, 24, 49, 24, 24, 24, 49, 24, 24, 24, 48, 24, 24, 24, 49, 24, 24, 24, 52, 24, 24, 24, 49, 24, 24, 24, 49, 24, 24, 24, 51, 24, 24, 24, 48, 24, 24, 24, 48, 24, 24, 24, 52, 24, 24, 24, 90, 23, 13, 24, 24, 24, 50, 24, 24, 24, 50, 24, 24, 24, 49, 24, 24, 24, 48, 24, 24, 24, 49, 24, 24, 24, 52, 24, 24, 24, 49, 24, 24, 24, 49, 24, 24, 24, 51, 24, 24, 24, 48, 24, 24, 24, 48, 24, 24, 24, 52, 24, 24, 24, 90, 24, 24, 24, 48, 24, 24, 24, 56, 24, 24, 24, 49, 11, 24, 24, 24, 48, 9, 6, 3, 24, 24, 24, 85, 4, 6, 19, 2, 24, 24, 24, 85, 24, 24, 24, 83, 24, 24, 24, 49, 22, 24, 24, 24, 48, 20, 6, 3, 24, 24, 24, 85, 4, 10, 12, 13, 24, 24, 24, 82, 24, 24, 24, 72, 24, 24, 24, 69, 24, 24, 24, 76, 24, 24, 24, 32, 24, 24, 24, 102, 24, 24, 24, 111, 24, 24, 24, 114, 24, 24, 24, 32, 24, 24, 24, 69, 24, 24, 24, 100, 24, 24, 24, 103, 24, 24, 24, 101, 24, 24, 24, 49, 17, 24, 24, 24, 48, 15, 6, 3, 24, 24, 24, 85, 4, 3, 12, 8, 24, 24, 24, 82, 24, 24, 24, 101, 24, 24, 24, 115, 24, 24, 24, 101, 24, 24, 24, 108, 24, 24, 24, 108, 24, 24, 24, 101, 24, 24, 24, 114, 24, 24, 24, 48, 24, 24, 24, 89, 24, 24, 24, 48, 19, 6, 7, 24, 24, 24, 42, 24, 24, 24, 134, 24, 24, 24, 72, 24, 24, 24, 206, 24, 24, 24, 61, 2, 1, 6, 8, 24, 24, 24, 42, 24, 24, 24, 134, 24, 24, 24, 72, 24, 24, 24, 206, 24, 24, 24, 61, 3, 1, 7, 3, 24, 24, 24, 66, 0, 4, 24, 24, 24, 34, 24, 24, 24, 45, 24, 24, 24, 210, 24, 24, 24, 218, 24, 24, 24, 250, 24, 24, 24, 41, 6, 24, 24, 24, 200, 24, 24, 24, 129, 24, 24, 24, 63, 24, 24, 24, 92, 24, 24, 24, 217, 24, 24, 24, 199, 8, 24, 24, 24, 183, 24, 24, 24, 47, 24, 24, 24, 63, 24, 24, 24, 36, 24, 24, 24, 136, 24, 24, 24, 211, 24, 24, 24, 31, 24, 24, 24, 100, 24, 24, 24, 135, 17, 24, 24, 24, 172, 24, 24, 24, 133, 24, 24, 24, 209, 24, 24, 24, 78, 24, 24, 24, 234, 24, 24, 24, 57, 24, 24, 24, 64, 24, 24, 24, 119, 24, 24, 24, 74, 24, 24, 24, 209, 24, 24, 24, 102, 24, 24, 24, 69, 24, 24, 24, 82, 24, 24, 24, 207, 24, 24, 24, 47, 24, 24, 24, 149, 18, 24, 24, 24, 60, 24, 24, 24, 116, 24, 24, 24, 24, 24, 24, 24, 60, 24, 24, 24, 153, 24, 24, 24, 111, 24, 24, 24, 79, 24, 24, 24, 137, 24, 24, 24, 84, 24, 24, 24, 40, 24, 24, 24, 118, 24, 24, 24, 157, 24, 24, 24, 220, 24, 24, 24, 206, 24, 24, 24, 99, 24, 24, 24, 176, 24, 24, 24, 222, 24, 24, 24, 168, 19, 24, 24, 24, 177, 24, 24, 24, 221, 24, 24, 24, 253, 24, 24, 24, 96, 24, 24, 24, 163, 24, 24, 24, 83, 24, 24, 24, 48, 24, 24, 24, 81, 24, 24, 24, 48, 24, 24, 24, 29, 6, 3, 24, 24, 24, 85, 24, 24, 24, 29, 14, 4, 22, 4, 20, 11, 24, 24, 24, 174, 13, 24, 24, 24, 198, 24, 24, 24, 133, 24, 24, 24, 152, 11, 24, 24, 24, 232, 24, 24, 24, 88, 24, 24, 24, 89, 24, 24, 24, 231, 24, 24, 24, 47, 24, 24, 24, 86, 24, 24, 24, 199, 24, 24, 24, 242, 24, 24, 24, 41, 24, 24, 24, 88, 8, 24, 24, 24, 135, 24, 24, 24, 182, 24, 24, 24, 48, 24, 24, 24, 31, 6, 3, 24, 24, 24, 85, 24, 24, 24, 29, 24, 24, 24, 35, 4, 24, 24, 24, 24, 24, 24, 24, 48, 22, 24, 24, 24, 128, 20, 11, 24, 24, 24, 174, 13, 24, 24, 24, 198, 24, 24, 24, 133, 24, 24, 24, 152, 11, 24, 24, 24, 232, 24, 24, 24, 88, 24, 24, 24, 89, 24, 24, 24, 231, 24, 24, 24, 47, 24, 24, 24, 86, 24, 24, 24, 199, 24, 24, 24, 242, 24, 24, 24, 41, 24, 24, 24, 88, 8, 24, 24, 24, 135, 24, 24, 24, 182, 24, 24, 24, 48, 15, 6, 3, 24, 24, 24, 85, 24, 24, 24, 29, 19, 1, 1, 24, 24, 24, 255, 4, 5, 24, 24, 24, 48, 3, 1, 1, 24, 24, 24, 255, 24, 24, 24, 48, 10, 6, 8, 24, 24, 24, 42, 24, 24, 24, 134, 24, 24, 24, 72, 24, 24, 24, 206, 24, 24, 24, 61, 4, 3, 2, 3, 24, 24, 24, 73, 0, 24, 24, 24, 48, 24, 24, 24, 70, 2, 24, 24, 24, 33, 0, 24, 24, 24, 154, 24, 24, 24, 151, 24, 24, 24, 24, 24, 24, 24, 121, 24, 24, 24, 134, 24, 24, 24, 120, 24, 24, 24, 42, 24, 24, 24, 99, 24, 24, 24, 186, 24, 24, 24, 73, 24, 24, 24, 211, 24, 24, 24, 241, 10, 24, 24, 24, 189, 24, 24, 24, 135, 24, 24, 24, 119, 24, 24, 24, 94, 24, 24, 24, 145, 20, 24, 24, 24, 72, 24, 24, 24, 33, 24, 24, 24, 139, 24, 24, 24, 81, 24, 24, 24, 136, 24, 24, 24, 90, 24, 24, 24, 162, 24, 24, 24, 89, 2, 24, 24, 24, 245, 24, 24, 24, 197, 24, 24, 24, 118, 24, 24, 24, 69, 2, 24, 24, 24, 33, 0, 24, 24, 24, 165, 24, 24, 24, 182, 24, 24, 24, 247, 17, 24, 24, 24, 174, 24, 24, 24, 184, 24, 24, 24, 106, 24, 24, 24, 223, 24, 24, 24, 139, 20, 24, 24, 24, 144, 21, 24, 24, 24, 43, 14, 24, 24, 24, 101, 24, 24, 24, 194, 24, 24, 24, 221, 24, 24, 24, 45, 24, 24, 24, 113, 24, 24, 24, 246, 24, 24, 24, 170, 24, 24, 24, 239, 24, 24, 24, 106, 24, 24, 24, 105, 16, 24, 24, 24, 145, 24, 24, 24, 180, 24, 24, 24, 178, 24, 24, 24, 117, 24, 24, 24, 51, 24, 24, 24, 191, 18, 24, 88, 24, 64, 24, 28, 24, 99, 17, 24, 88, 13, 24, 39, 24, 231, 24, 65, 24, 152, 24, 43, 24, 200, 24, 155, 24, 77, 24, 230, 24, 40, 24, 251, 24, 247, 24, 122, 24, 41, 24, 128, 24, 53, 24, 75, 24, 26, 24, 119, 13, 24, 217, 24, 86, 24, 178, 24, 159, 21, 24, 171, 24, 212, 24, 44, 24, 238, 24, 127, 24, 103, 24, 143, 24, 136, 2, 24, 50, 24, 198, 24, 52, 24, 215, 24, 35, 24, 98, 24, 230, 24, 173, 24, 218, 24, 146, 24, 169, 24, 181, 24, 84, 24, 228, 24, 46, 24, 43, 24, 246, 24, 212, 24, 58, 24, 54, 24, 181, 24, 249, 24, 33, 24, 99, 24, 240, 153, 4, 24, 24, 210, 24, 132, 24, 67, 24, 161, 1, 24, 38, 24, 160, 24, 89, 3, 24, 204, 24, 131, 24, 130, 24, 56, 24, 42, 24, 152, 24, 48, 24, 24, 24, 171, 24, 24, 24, 254, 24, 24, 24, 150, 24, 24, 24, 32, 24, 24, 24, 79, 24, 24, 24, 232, 24, 24, 24, 71, 24, 24, 24, 112, 24, 24, 24, 207, 24, 24, 24, 196, 24, 24, 24, 61, 24, 24, 24, 82, 24, 24, 24, 143, 24, 24, 24, 230, 24, 24, 24, 118, 24, 24, 24, 246, 24, 24, 24, 98, 21, 24, 24, 24, 84, 24, 24, 24, 140, 24, 24, 24, 179, 24, 24, 24, 162, 24, 24, 24, 53, 8, 24, 24, 24, 173, 3, 5, 11, 24, 24, 24, 187, 24, 24, 24, 108, 24, 24, 24, 64, 24, 24, 24, 204, 24, 24, 24, 179, 9, 24, 24, 24, 127, 24, 24, 24, 132, 12, 24, 24, 24, 88, 24, 24, 24, 234, 24, 24, 24, 116, 24, 24, 24, 46, 24, 24, 24, 58, 24, 24, 24, 64, 24, 24, 24, 54, 24, 24, 24, 85, 24, 24, 24, 53, 24, 24, 24, 99, 24, 24, 24, 54, 24, 130, 24, 56, 24, 42, 24, 152, 24, 48, 24, 24, 24, 100, 24, 24, 24, 109, 24, 24, 24, 185, 24, 24, 24, 81, 24, 24, 24, 179, 24, 24, 24, 99, 24, 24, 24, 166, 24, 24, 24, 229, 12, 24, 24, 24, 227, 24, 24, 24, 245, 24, 24, 24, 145, 6, 24, 24, 24, 54, 24, 24, 24, 103, 24, 24, 24, 44, 24, 24, 24, 29, 24, 24, 24, 185, 24, 24, 24, 168, 24, 24, 24, 33, 24, 24, 24, 120, 24, 24, 24, 236, 24, 24, 24, 117, 24, 24, 24, 252, 19, 24, 24, 24, 253, 24, 24, 24, 154, 24, 24, 24, 235, 24, 24, 24, 70, 24, 24, 24, 89, 24, 24, 24, 151, 24, 24, 24, 159, 24, 24, 24, 241, 24, 24, 24, 247, 24, 24, 24, 43, 24, 24, 24, 243, 24, 24, 24, 230, 24, 24, 24, 130, 24, 24, 24, 80, 24, 24, 24, 126, 24, 24, 24, 101, 24, 24, 24, 86, 24, 24, 24, 35, 21, 24, 24, 24, 31, 24, 24, 24, 175, 24, 24, 24, 110, 24, 24, 24, 137, 24, 131, 24, 38, 1, 24, 153, 1, 24, 195, 24, 24, 24, 48, 24, 24, 24, 130, 1, 24, 24, 24, 191, 24, 24, 24, 48, 24, 24, 24, 130, 1, 24, 24, 24, 101, 24, 24, 24, 160, 3, 2, 1, 2, 2, 20, 24, 24, 24, 98, 24, 24, 24, 124, 24, 24, 24, 85, 24, 24, 24, 60, 24, 24, 24, 78, 24, 24, 24, 145, 24, 24, 24, 128, 24, 24, 24, 247, 24, 24, 24, 40, 4, 24, 24, 24, 51, 24, 24, 24, 199, 24, 24, 24, 165, 24, 24, 24, 221, 24, 24, 24, 80, 24, 24, 24, 80, 24, 24, 24, 250, 24, 24, 24, 41, 24, 24, 24, 235, 24, 24, 24, 60, 24, 24, 24, 48, 10, 6, 8, 24, 24, 24, 42, 24, 24, 24, 134, 24, 24, 24, 72, 24, 24, 24, 206, 24, 24, 24, 61, 4, 3, 2, 24, 24, 24, 48, 24, 24, 24, 53, 24, 24, 24, 49, 11, 24, 24, 24, 48, 9, 6, 3, 24, 24, 24, 85, 4, 6, 19, 2, 24, 24, 24, 85, 24, 24, 24, 83, 24, 24, 24, 49, 22, 24, 24, 24, 48, 20, 6, 3, 24, 24, 24, 85, 4, 10, 12, 13, 24, 24, 24, 82, 24, 24, 24, 72, 24, 24, 24, 69, 24, 24, 24, 76, 24, 24, 24, 32, 24, 24, 24, 102, 24, 24, 24, 111, 24, 24, 24, 114, 24, 24, 24, 32, 24, 24, 24, 69, 24, 24, 24, 100, 24, 24, 24, 103, 24, 24, 24, 101, 24, 24, 24, 49, 14, 24, 24, 24, 48, 12, 6, 3, 24, 24, 24, 85, 4, 3, 12, 5, 24, 24, 24, 79, 24, 24, 24, 119, 24, 24, 24, 110, 24, 24, 24, 101, 24, 24, 24, 114, 24, 24, 24, 48, 24, 24, 24, 30, 23, 13, 24, 24, 24, 50, 24, 24, 24, 49, 24, 24, 24, 49, 24, 24, 24, 48, 24, 24, 24, 49, 24, 24, 24, 52, 24, 24, 24, 49, 24, 24, 24, 49, 24, 24, 24, 51, 24, 24, 24, 48, 24, 24, 24, 48, 24, 24, 24, 52, 24, 24, 24, 90, 23, 13, 24, 24, 24, 50, 24, 24, 24, 50, 24, 24, 24, 49, 24, 24, 24, 48, 24, 24, 24, 49, 24, 24, 24, 52, 24, 24, 24, 49, 24, 24, 24, 49, 24, 24, 24, 51, 24, 24, 24, 48, 24, 24, 24, 48, 24, 24, 24, 52, 24, 24, 24, 90, 24, 24, 24, 48, 24, 24, 24, 53, 24, 24, 24, 49, 11, 24, 24, 24, 48, 9, 6, 3, 24, 24, 24, 85, 4, 6, 19, 2, 24, 24, 24, 85, 24, 24, 24, 83, 24, 24, 24, 49, 22, 24, 24, 24, 48, 20, 6, 3, 24, 24, 24, 85, 4, 10, 12, 13, 24, 24, 24, 82, 24, 24, 24, 72, 24, 24, 24, 69, 24, 24, 24, 76, 24, 24, 24, 32, 24, 24, 24, 102, 24, 24, 24, 111, 24, 24, 24, 114, 24, 24, 24, 32, 24, 24, 24, 69, 24, 24, 24, 100, 24, 24, 24, 103, 24, 24, 24, 101, 24, 24, 24, 49, 14, 24, 24, 24, 48, 12, 6, 3, 24, 24, 24, 85, 4, 3, 12, 5, 24, 24, 24, 79, 24, 24, 24, 119, 24, 24, 24, 110, 24, 24, 24, 101, 24, 24, 24, 114, 24, 24, 24, 48, 24, 24, 24, 89, 24, 24, 24, 48, 19, 6, 7, 24, 24, 24, 42, 24, 24, 24, 134, 24, 24, 24, 72, 24, 24, 24, 206, 24, 24, 24, 61, 2, 1, 6, 8, 24, 24, 24, 42, 24, 24, 24, 134, 24, 24, 24, 72, 24, 24, 24, 206, 24, 24, 24, 61, 3, 1, 7, 3, 24, 24, 24, 66, 0, 4, 24, 24, 24, 195, 24, 24, 24, 174, 24, 24, 24, 63, 24, 24, 24, 182, 24, 24, 24, 65, 24, 24, 24, 104, 24, 24, 24, 35, 24, 24, 24, 26, 24, 24, 24, 193, 24, 24, 24, 154, 24, 24, 24, 146, 24, 24, 24, 119, 24, 24, 24, 66, 24, 24, 24, 45, 24, 24, 24, 58, 24, 24, 24, 185, 24, 24, 24, 145, 24, 24, 24, 72, 24, 24, 24, 162, 24, 24, 24, 182, 24, 24, 24, 107, 24, 24, 24, 31, 24, 24, 24, 247, 24, 24, 24, 94, 24, 24, 24, 64, 24, 24, 24, 96, 24, 24, 24, 206, 24, 24, 24, 187, 24, 24, 24, 71, 24, 24, 24, 137, 24, 24, 24, 108, 24, 24, 24, 235, 24, 24, 24, 140, 24, 24, 24, 38, 24, 24, 24, 138, 24, 24, 24, 131, 24, 24, 24, 62, 24, 24, 24, 235, 24, 24, 24, 212, 24, 24, 24, 139, 24, 24, 24, 193, 24, 24, 24, 240, 24, 24, 24, 133, 24, 24, 24, 202, 24, 24, 24, 214, 24, 24, 24, 31, 24, 24, 24, 61, 24, 24, 24, 147, 24, 24, 24, 114, 24, 24, 24, 170, 24, 24, 24, 225, 24, 24, 24, 250, 24, 24, 24, 179, 24, 24, 24, 177, 24, 24, 24, 134, 24, 24, 24, 174, 24, 24, 24, 45, 24, 24, 24, 79, 16, 20, 24, 24, 24, 42, 24, 24, 24, 116, 24, 24, 24, 207, 24, 24, 24, 31, 24, 24, 24, 163, 24, 24, 24, 83, 24, 24, 24, 48, 24, 24, 24, 81, 24, 24, 24, 48, 24, 24, 24, 29, 6, 3, 24, 24, 24, 85, 24, 24, 24, 29, 14, 4, 22, 4, 20, 8, 24, 24, 24, 157, 24, 24, 24, 253, 24, 24, 24, 44, 24, 24, 24, 223, 24, 24, 24, 201, 24, 24, 24, 51, 24, 24, 24, 235, 24, 24, 24, 220, 24, 24, 24, 205, 24, 24, 24, 169, 24, 24, 24, 35, 24, 24, 24, 59, 24, 24, 24, 172, 24, 24, 24, 114, 24, 24, 24, 205, 24, 24, 24, 187, 24, 24, 24, 143, 24, 24, 24, 146, 24, 24, 24, 88, 24, 24, 24, 48, 24, 24, 24, 31, 6, 3, 24, 24, 24, 85, 24, 24, 24, 29, 24, 24, 24, 35, 4, 24, 24, 24, 24, 24, 24, 24, 48, 22, 24, 24, 24, 128, 20, 8, 24, 24, 24, 157, 24, 24, 24, 253, 24, 24, 24, 44, 24, 24, 24, 223, 24, 24, 24, 201, 24, 24, 24, 51, 24, 24, 24, 235, 24, 24, 24, 220, 24, 24, 24, 205, 24, 24, 24, 169, 24, 24, 24, 35, 24, 24, 24, 59, 24, 24, 24, 172, 24, 24, 24, 114, 24, 24, 24, 205, 24, 24, 24, 187, 24, 24, 24, 143, 24, 24, 24, 146, 24, 24, 24, 88, 24, 24, 24, 48, 15, 6, 3, 24, 24, 24, 85, 24, 24, 24, 29, 19, 1, 1, 24, 24, 24, 255, 4, 5, 24, 24, 24, 48, 3, 1, 1, 24, 24, 24, 255, 24, 24, 24, 48, 10, 6, 8, 24, 24, 24, 42, 24, 24, 24, 134, 24, 24, 24, 72, 24, 24, 24, 206, 24, 24, 24, 61, 4, 3, 2, 3, 24, 24, 24, 72, 0, 24, 24, 24, 48, 24, 24, 24, 69, 2, 24, 24, 24, 33, 0, 24, 24, 24, 157, 24, 24, 24, 124, 11, 24, 24, 24, 238, 24, 24, 24, 53, 3, 24, 24, 24, 81, 24, 24, 24, 64, 24, 24, 24, 211, 8, 24, 24, 24, 197, 24, 24, 24, 223, 8, 24, 24, 24, 65, 2, 24, 24, 24, 160, 24, 24, 24, 51, 24, 24, 24, 105, 24, 24, 24, 81, 24, 24, 24, 66, 24, 24, 24, 220, 24, 24, 24, 102, 24, 24, 24, 193, 24, 24, 24, 156, 24, 24, 24, 93, 24, 24, 24, 168, 24, 24, 24, 172, 24, 24, 24, 77, 24, 24, 24, 118, 6, 24, 24, 24, 106, 4, 2, 24, 24, 24, 32, 24, 24, 24, 98, 24, 24, 24, 137, 24, 24, 24, 180, 24, 24, 24, 132, 24, 24, 24, 162, 24, 24, 24, 131, 24, 24, 24, 247, 24, 24, 24, 217, 12, 24, 24, 24, 251, 24, 24, 24, 238, 24, 24, 24, 78, 24, 24, 24, 201, 24, 24, 24, 195, 23, 0, 24, 24, 24, 125, 24, 24, 24, 72, 24, 24, 24, 201, 24, 24, 24, 115, 24, 24, 24, 197, 24, 24, 24, 159, 24, 24, 24, 220, 24, 24, 24, 130, 24, 24, 24, 247, 24, 24, 24, 228, 24, 24, 24, 139, 24, 24, 24, 57, 24, 24, 24, 112, 24, 24, 24, 43, 24, 24, 24, 98, 24, 24, 24, 215, 24, 88, 24, 64, 24, 88, 24, 157, 24, 158, 24, 161, 24, 219, 24, 176, 24, 71, 24, 91, 17, 21, 24, 253, 13, 24, 25, 24, 212, 24, 203, 24, 197, 24, 252, 24, 95, 24, 211, 24, 43, 19, 24, 210, 3, 24, 164, 24, 157, 24, 27, 24, 83, 24, 49, 24, 124, 24, 214, 24, 136, 24, 67, 24, 166, 17, 24, 241, 24, 187, 24, 93, 24, 56, 24, 64, 24, 216, 24, 40, 24, 129, 24, 194, 24, 73, 24, 34, 24, 169, 24, 241, 24, 176, 24, 107, 24, 92, 24, 60, 24, 171, 24, 252, 24, 126, 24, 85, 24, 37, 24, 148, 24, 196, 24, 236, 24, 97, 24, 152, 24, 61, 24, 253, 24, 71}
)

var _ = Describe("Ownershipvoucher", func() {
	Describe("extract minimum data required", func() {
		Context("parse", func() {
			data, e := ovde.MinimumParse(ovb)
			It("no error", func ()  {
				Expect(e).ToNot(HaveOccurred())
			})
			It("has correct data", func ()  {
				Expect(len(data)).To(Equal(1))
				Expect(data[0]["device_name"]).To(Equal(edgeDeviceName))
				Expect(data[0]["fdo_uuid"]).To(Equal(edgeDeviceUUID))
				Expect(data[0]["protocol_version"]).To(Equal(uint16(100)))
			})
		})
	})

	Describe("break guid", func() {
		Context("parse the bytes only", func() {
			mParse, e := ovde.ParseBytes(ovb)
			When("parsing", func() {
				It("error", func() {
					Expect(e).To(HaveOccurred())
				})
			})
			ovh := mParse[0]
			ovh.GUID = []byte{0}
			b, e := ovh.MarshalJSON()
			It("error", func() {
				Expect(e).To(HaveOccurred())
			})
			It("contains invalid data", func ()  {
				j := map[string]interface{}{}
				e := json.Unmarshal(b, &j)
				By("no error", func() {
					Expect(e).ToNot(HaveOccurred())
				})
				By("json has correct values", func() {
					Expect(j["device_info"]).To(Equal(edgeDeviceName))
					Expect(j["guid"]).To(Equal("null"))
					Expect(j["protocol_version"]).To(Equal(float64(100)))
				}) 
			})
		})
	})


	Describe("ownershipvoucher header marshal JSON", func() {
		Context("parse the bytes only", func() {
			mParse, e := ovde.ParseBytes(ovb)
			When("parsing", func() {
				It("no error", func() {
					Expect(e).ToNot(HaveOccurred())
				})
			})
			ovh := mParse[0]
			b, e := ovh.MarshalJSON()
			When("marshal JSON", func() {
				It("no error", func() {
					Expect(e).ToNot(HaveOccurred())
				})
				It("not empty", func() {
					Expect(b).ToNot(BeEmpty())
				})
				It("contains valid data", func ()  {
					j := map[string]interface{}{}
					e := json.Unmarshal(b, &j)
					By("no error", func() {
						Expect(e).ToNot(HaveOccurred())
					})
					By("json has correct values", func() {
						Expect(j["device_info"]).To(Equal(edgeDeviceName))
						Expect(j["guid"]).To(Equal(edgeDeviceUUID))
						Expect(j["protocol_version"]).To(Equal(float64(100)))
					}) 
				})
			})
		})
	})

	Describe("check resolvers", func() {
		Context("ResolvePublicKeyEncoding testing", func() {
			It("should succeed", func() {
				for i := range [5]int{} {
					switch enc, e := models.ResolvePublicKeyEncoding(i); enc {
					case "Crypto":
						Expect(i).To(Equal(0))
					case "X509":
						Expect(i).To(Equal(1))
					case "COSEX509":
						Expect(i).To(Equal(2))
					case "Cosekey":
						Expect(i).To(Equal(3))
					case "":
						Expect(e).To(HaveOccurred())
					}
				}
			})
		})
		Context("ResolveRendezvousVariableCode testing", func() {
			It("should succeed", func() {
				for i := range [17]int{} {
					switch v, e:= models.ResolveRendezvousVariableCode(i);v {
					case "DeviceOnly":
						Expect(i).To(Equal(0))
					case "OwnerOnly":
						Expect(i).To(Equal(1))
					case "IPAddress":
						Expect(i).To(Equal(2))
					case "DevicePort":
						Expect(i).To(Equal(3))
					case "OwnerPort":
						Expect(i).To(Equal(4))
					case "Dns":
						Expect(i).To(Equal(5))
					case "ServerCertHash":
						Expect(i).To(Equal(6))
					case "CaCertHash":
						Expect(i).To(Equal(7))
					case "UserInput":
						Expect(i).To(Equal(8))
					case "WifiSsid":
						Expect(i).To(Equal(9))
					case "WifiPw":
						Expect(i).To(Equal(10))
					case "Medium":
						Expect(i).To(Equal(11))
					case "Protocol":
						Expect(i).To(Equal(12))
					case "Delaysec":
						Expect(i).To(Equal(13))
					case "Bypass":
						Expect(i).To(Equal(14))
					case "Extended":
						Expect(i).To(Equal(15))
					case "":
						Expect(e).To(HaveOccurred())
					}
				}
			})
		})
	})
})
