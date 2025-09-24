GoLog Analyzer – Analyse de Logs Distribuée
📌 Contexte

Ce projet est un outil en ligne de commande (CLI) écrit en Go, nommé loganalyzer.
Il permet aux administrateurs système d’analyser plusieurs fichiers de logs en parallèle, d’en extraire un résumé clair et d’exporter les résultats au format JSON.

L’objectif est de centraliser l’analyse de fichiers provenant de différentes sources (serveurs, applications), tout en gérant les erreurs de manière robuste.

🎯 Objectifs pédagogiques

Ce projet met en pratique plusieurs concepts clés du langage Go :

Concurrence : utilisation de goroutines et WaitGroups.

Gestion des erreurs : création d’erreurs personnalisées et usage de errors.Is / errors.As.

Développement CLI : structuration avec Cobra et gestion de drapeaux (flags).

Import/Export JSON : lecture d’un fichier de configuration et génération d’un rapport.

Modularité : séparation du code en packages logiques.

⚙️ Installation & Prérequis

Avoir Go installé (version ≥ 1.20).

Cloner le projet :

git clone https://github.com/votre-utilisateur/loganalyzer.git
cd loganalyzer


Initialiser les dépendances :

go mod tidy

🚀 Utilisation
Lancer une analyse
go run . analyze -c config.json -o report.json


🗂️ Architecture du projet
.
├── cmd/                 # Commandes CLI
│   ├── root.go          # Commande racine
│   └── analyze.go       # Commande analyze
├── internal/
│   ├── analyzer/        # Logique d’analyse + erreurs personnalisées
│   │   ├── analyzer.go
│   │   └── errors.go
│   ├── config/          # Chargement du fichier config JSON
│   │   └── config.go
│   └── reporter/        # Export JSON
│       └── exporter.go
├── test_logs/           # Exemples de fichiers de logs
├── config.json          # Fichier de configuration
├── go.mod               # Dépendances Go
└── main.go              # Point d’entrée

👥 Équipe

- Lucas Labeye
