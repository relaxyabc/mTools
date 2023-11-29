package cmd

import (
	"bufio"
	"errors"
	"fmt"
	sum "github.com/toolbox/hash"
	"github.com/urfave/cli/v2"
	"io"
	"os"
)

// Algorithms MD2 MD4 MD5 SHA1 SHA256 SHA384 SHA512
var Algorithms = []string{"crc32", "md5", "sha1", "sha256", "sha384", "sha512"}

var Hashes = make(map[string]sum.Hash)
var AlgorithmsImpl []sum.Hash

func init() {
	AlgorithmsImpl = append(AlgorithmsImpl, sum.NewMd5())
	AlgorithmsImpl = append(AlgorithmsImpl, sum.NewSha1())
	AlgorithmsImpl = append(AlgorithmsImpl, sum.NewCrc32())
	AlgorithmsImpl = append(AlgorithmsImpl, sum.NewSha256())
	AlgorithmsImpl = append(AlgorithmsImpl, sum.NewSha384())
	AlgorithmsImpl = append(AlgorithmsImpl, sum.NewSha512())

	for _, item := range AlgorithmsImpl {
		Hashes[item.GetAlgorithmName()] = item
	}
}

// NewHashCmd hashCommand
func NewHashCmd() *cli.Command {
	return &cli.Command{
		Name:    "summary",
		Aliases: []string{"sum", "s"},
		Usage:   "toolbox sum",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "algorithm",
				Aliases: []string{"a"},
				Value:   "all",
				Usage:   "all/crc32/md5/sha1/sha256/sha384/sha512",
				Action:  algorithmCheck(),
			},
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "toolbox summary -a md5 -f /home/user/abc.txt",
				Action:  fileExistsCheck(),
			},
			&cli.StringFlag{
				Name:    "content",
				Aliases: []string{"c"},
				Usage:   "toolbox summary -a md5 -c \"content\"",
			},
		},
		Action: func(context *cli.Context) error {
			a := context.Value("algorithm")
			algorithm := a.(string)
			f := context.Value("file")
			file := f.(string)
			c := context.Value("content")
			content := c.(string)
			if file == "" && content == "" {
				return errors.New("file and content is all empty")
			}

			if file != "" {
				switch algorithm {
				case "all":
					err := calculateFileHash(file)
					if err != nil {
						return err
					}
				case "md5", "crc32", "sha1", "sha256", "sha384", "sha512":
					err := calculateFileHashByAlgorithm(file, algorithm)
					if err != nil {
						return err
					}
				default:
					return errors.New("unknown algorithm: " + algorithm)
				}
			}
			if content != "" {
				switch algorithm {
				case "all":
					err := calculateContentHash(content)
					if err != nil {
						return err
					}
				case "md5", "crc32", "sha1", "sha256", "sha384", "sha512":
					err := calculateContentHashByAlgorithm(content, algorithm)
					if err != nil {
						return err
					}
				default:
					return errors.New("unknown algorithm: " + algorithm)
				}
			}

			return nil
		},
	}
}

func calculateFileHash(file string) error {
	of, err := os.Open(file)
	if err != nil {
		return err
	}
	defer func(of *os.File) {
		err := of.Close()
		if err != nil {

		}
	}(of)

	for buf, reader := make([]byte, 1<<16), bufio.NewReader(of); ; {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		for _, v := range AlgorithmsImpl {
			v.WriteContent(buf[:n])
		}
	}
	fmt.Printf("algorithm\t | hash值 \t\n")
	for _, v := range AlgorithmsImpl {
		hashVal := v.CalculateSum()
		fmt.Printf("%s\t\t | %s \t\n", v.GetAlgorithmName(), hashVal)
	}
	return nil
}

func calculateFileHashByAlgorithm(file string, algorithm string) error {
	of, err := os.Open(file)
	if err != nil {
		return err
	}
	defer func(of *os.File) {
		err := of.Close()
		if err != nil {

		}
	}(of)

	hash := Hashes[algorithm]
	for buf, reader := make([]byte, 1<<16), bufio.NewReader(of); ; {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		hash.WriteContent(buf[:n])

	}
	fmt.Printf("algorithm\t | hash值 \t\n")
	hashVal := hash.CalculateSum()
	fmt.Printf("%s\t\t | %s \t\n", algorithm, hashVal)
	return nil
}

func calculateContentHash(content string) error {
	fmt.Printf("algorithm\t | hash值 \t\n")
	bytes := []byte(content)
	for _, v := range AlgorithmsImpl {
		v.WriteContent(bytes)
		hashVal := v.CalculateSum()
		fmt.Printf("%s\t\t | %s \t\n", v.GetAlgorithmName(), hashVal)
	}
	return nil
}

func calculateContentHashByAlgorithm(content string, algorithm string) error {
	fmt.Printf("algorithm\t | hash值 \t\n")
	bytes := []byte(content)
	hash := Hashes[algorithm]
	hash.WriteContent(bytes)
	hashVal := hash.CalculateSum()
	fmt.Printf("%s\t\t | %s \t\n", algorithm, hashVal)
	return nil
}

// fileExistsCheck 文件是否存在检查
func fileExistsCheck() func(context *cli.Context, file string) error {
	return func(context *cli.Context, file string) error {
		if file == "" {
			return nil
		}
		_, err := os.Stat(file)
		if err != nil {
			return err
		}

		return nil
	}
}

// algorithmCheck 算法参数校验
func algorithmCheck() func(context *cli.Context, algorithm string) error {
	return func(context *cli.Context, algorithm string) error {
		if algorithm == "" || algorithm == "all" {
			return nil
		}
		var checked = false
		for _, item := range Algorithms {
			if item == algorithm {
				checked = true
				break
			}
		}
		if !checked {
			return errors.New("illegal algorithm: " + algorithm)
		}
		return nil
	}
}
